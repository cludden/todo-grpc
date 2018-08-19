package resolvers

import (
	"context"
	"strconv"
	"time"
	"todo-grpc/graphql/models"
	"todo-grpc/proto"

	"github.com/golang/protobuf/ptypes"
	"github.com/vektah/gqlparser/gqlerror"
	null "gopkg.in/guregu/null.v3"
)

// Todos implements the Todo type resolver
type Todos Resolver

// CompletedAt field resolver
func (t *Todos) CompletedAt(ctx context.Context, todo *models.Todo) (*time.Time, error) {
	if todo == nil {
		return nil, nil
	}
	return todo.CompletedAt, nil
}

// CreatedAt field resolver
func (t *Todos) CreatedAt(ctx context.Context, todo *models.Todo) (time.Time, error) {
	if todo == nil {
		return time.Time{}, nil
	}
	return todo.CreatedAt, nil
}

// CompleteTodo marks an existing Todo as complete
func (m *Mutations) CompleteTodo(ctx context.Context, input models.CompleteTodoInput) (*models.Todo, error) {
	res, err := m.todos.CompleteTodo(ctx, &proto.CompleteTodoInput{
		Id: input.ID,
	})
	if err != nil {
		return nil, gqlerror.Errorf("todos error: %v", err)
	}

	return unmarshalTodoPB(res.GetTodo())
}

// CreateTodo creates a new Todo
func (m *Mutations) CreateTodo(ctx context.Context, input models.CreateTodoInput) (*models.Todo, error) {
	res, err := m.todos.CreateTodo(ctx, &proto.CreateTodoInput{
		Description: null.StringFromPtr(input.Description).ValueOrZero(),
		Title:       input.Title,
	})
	if err != nil {
		return nil, gqlerror.Errorf("todos error: %v", err)
	}

	return unmarshalTodoPB(res.GetTodo())
}

// Todos retrieves a paginated list of Todos
func (q *Query) Todos(ctx context.Context, input models.TodosQueryInput) (*models.TodosQueryConnection, error) {
	// define query params
	params := proto.ListTodosInput{
		After: null.StringFromPtr(input.After).ValueOrZero(),
		First: 20,
		Query: null.StringFromPtr(input.Query).ValueOrZero(),
	}
	if input.First != nil {
		params.First = int32(*input.First)
	}

	// verify pagination cursor
	var after int64
	if params.After != "" {
		a, err := strconv.ParseInt(params.After, 10, 64)
		if err != nil {
			return nil, gqlerror.Errorf("invalid cursor: %s", params.After)
		}
		after = a
	}

	// execute service query
	res, err := q.todos.ListTodos(ctx, &params)
	if err != nil {
		return nil, gqlerror.Errorf("todos error: %v", err)
	}

	// build result
	n, todos := int(res.GetTotal()), res.GetTodos()
	out := models.TodosQueryConnection{
		Edges: make([]*models.TodosQueryEdge, len(todos)),
		PageInfo: &models.PageInfo{
			Total: &n,
		},
	}
	for i, t := range res.GetTodos() {
		todo, err := unmarshalTodoPB(t)
		if err != nil {
			return nil, err
		}
		after++
		out.Edges[i] = &models.TodosQueryEdge{
			Cursor: strconv.FormatInt(after, 10),
			Node:   todo,
		}
	}

	return &out, nil
}

// unmarshalTodoPB unmarshals a Todo protobuf message into a graphql type
func unmarshalTodoPB(t *proto.Todo) (*models.Todo, error) {
	todo := models.Todo{
		ID:          t.GetId(),
		Complete:    t.GetComplete(),
		Description: null.StringFrom(t.GetDescription()).Ptr(),
		Title:       t.GetTitle(),
	}

	completedAt, err := ptypes.Timestamp(t.GetCompletedAt())
	valid := true
	if err != nil {
		valid = false
	}
	todo.CompletedAt = null.NewTime(completedAt, valid).Ptr()

	createdAt, err := ptypes.Timestamp(t.GetCreatedAt())
	if err != nil {
		return nil, gqlerror.Errorf("error marshalling todo: %v", err)
	}
	todo.CreatedAt = createdAt

	return &todo, nil
}
