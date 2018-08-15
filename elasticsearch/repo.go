// Package elasticsearch provides types and methods for interacting with an
// ElasticSearch todos store
package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"todo-grpc/proto"
	"todo-grpc/todo"
	"todo-grpc/validation"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

// Config describes the input to a NewRepository operation
type Config struct {
	Index    string             `validate:"required"`
	Log      logrus.FieldLogger `validate:"required"`
	URL      string             `validate:"required"`
	TypeName string             `validate:"required"`
}

// Repository provides access to an ElasticSearch backed todos store
type Repository struct {
	client   *elastic.Client
	index    string
	log      logrus.FieldLogger
	typeName string
}

// NewRepository returns a new repository value
func NewRepository(c *Config) (*Repository, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// create elasticsearch client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(c.URL),
		elastic.SetHealthcheckTimeout(time.Second*30),
		elastic.SetErrorLog(NewLog(c.Log.Errorf)),
		elastic.SetInfoLog(NewLog(c.Log.Infof)),
		elastic.SetTraceLog(NewLog(c.Log.Debugf)),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating elasticsearch client: %v", err)
	}

	r := Repository{
		client:   client,
		index:    c.Index,
		log:      c.Log,
		typeName: c.TypeName,
	}
	return &r, nil
}

// CompleteTodo markes an existing todo as completed
func (r *Repository) CompleteTodo(ctx context.Context, id string) (*todo.Todo, error) {
	// define update script
	script := elastic.NewScript("ctx._source.complete = params.complete; ctx._source.completed_at = params.completed_at").
		Lang("painless").
		Type("inline").
		Param("complete", true).
		Param("completed_at", time.Now())

	// define update operation
	update := r.client.Update().
		Index(r.index).
		Type(r.typeName).
		Id(id).
		Script(script).
		FetchSource(true)

	// execute index operation
	res, err := update.Do(ctx)
	if err != nil {
		r.log.WithError(err).Errorln("elasticsearch:error:create-todo")
		return nil, err
	}

	// unmarshal output
	if res.GetResult == nil || res.GetResult.Source == nil {
		return nil, errors.New("not found")
	}
	var t todo.Todo
	if err := json.Unmarshal(*res.GetResult.Source, &t); err != nil {
		return nil, fmt.Errorf("error unmarshaling todo: %v", err)
	}
	return &t, nil
}

// CreateTodo creates a new todo
func (r *Repository) CreateTodo(ctx context.Context, in *todo.Todo) error {
	// define index operation
	index := r.client.Index().
		Index(r.index).
		Type(r.typeName).
		Id(in.ID).
		Refresh("true").
		BodyJson(in)

	// execute index operation
	_, err := index.Do(ctx)
	if err != nil {
		r.log.WithError(err).Errorln("elasticsearch:error:create-todo")
		return err
	}
	return nil
}

// ListTodos retrieves a paginated list of todos
func (r *Repository) ListTodos(ctx context.Context, in *proto.ListTodosInput) (*todo.ListTodosOutput, error) {
	var out todo.ListTodosOutput

	// define base query
	var query elastic.Query
	if q := in.GetQuery(); q != "" {
		query = elastic.NewSimpleQueryStringQuery(q)
	} else {
		query = elastic.NewMatchAllQuery()
	}

	// define base search operation
	search := r.client.Search(r.index).
		Query(query).
		Size(int(in.GetFirst())).
		Sort("created_at", false)

	// add pagination offset if applicable
	if after := in.GetAfter(); after != "" {
		from, err := strconv.Atoi(after)
		if err != nil || from < 0 {
			return nil, fmt.Errorf("invalid pagination offset: %v", err)
		}
		search = search.From(from)
	}

	// execute search operation
	res, err := search.Do(ctx)
	if err != nil {
		if eserr, ok := err.(*elastic.Error); ok && eserr.Details.Type == "index_not_found_exception" {
			fmt.Printf("%+v\n", eserr.Details)
			return &out, nil
		}
		r.log.WithError(err).Errorln("elasticsearch:error:list-todos")
		return nil, err
	}

	// build output
	out.Total = int32(res.TotalHits())
	out.Todos = make([]*todo.Todo, len(res.Hits.Hits))
	for i, d := range res.Hits.Hits {
		var t todo.Todo
		if d.Source == nil {
			return &out, fmt.Errorf("missing todo source for document with id: %s", d.Id)
		}
		if err := json.Unmarshal(*d.Source, &t); err != nil {
			r.log.WithError(err).Errorln("elasticsearch:error:list-todos:unmarsha")
			return &out, fmt.Errorf("unable to unmarshal todo: %v", err)
		}
		out.Todos[i] = &t
	}

	return &out, nil
}
