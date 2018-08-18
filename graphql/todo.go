package graphql

import (
	"context"
	"time"
	"todo-grpc/graphql/models"
	"todo-grpc/proto"

	"github.com/golang/protobuf/ptypes"

	null "gopkg.in/guregu/null.v3"
)

// Todos implements the Todo type resolver
type Todos struct {
}

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
		return nil, err
	}

	return UnmarshalTodoPB(res.GetTodo())
}

// CreateTodo creates a new Todo
func (m *Mutations) CreateTodo(ctx context.Context, input models.CreateTodoInput) (*models.Todo, error) {
	res, err := m.todos.CreateTodo(ctx, &proto.CreateTodoInput{
		Description: null.StringFromPtr(input.Description).ValueOrZero(),
		Title:       input.Title,
	})
	if err != nil {
		return nil, err
	}

	return UnmarshalTodoPB(res.GetTodo())
}

// Todos retrieves a paginated list of Todos
func (q *Query) Todos(ctx context.Context, input models.TodosQueryInput) (*models.TodosQueryConnection, error) {
	params := proto.ListTodosInput{
		After: null.StringFromPtr(input.After).ValueOrZero(),
		Query: null.StringFromPtr(input.Query).ValueOrZero(),
	}
	if input.First != nil {
		params.First = int32(*input.First)
	}

	res, err := q.todos.ListTodos(ctx, &params)
	if err != nil {
		return nil, err
	}

	n := int(res.GetTotal())
	out := models.TodosQueryConnection{
		Edges: make([]*models.TodosQueryEdge, n),
		PageInfo: &models.PageInfo{
			Total: &n,
		},
	}
	for i, t := range res.GetTodos() {
		todo, err := UnmarshalTodoPB(t)
		if err != nil {
			return nil, err
		}
		out.Edges[i] = &models.TodosQueryEdge{
			Cursor: todo.ID,
			Node:   todo,
		}
	}

	return &out, nil
}

// UnmarshalTodoPB unmarshals a Todo protobuf message into a graphql type
func UnmarshalTodoPB(t *proto.Todo) (*models.Todo, error) {
	todo := models.Todo{
		ID:          t.GetTitle(),
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
		return nil, err
	}
	todo.CreatedAt = createdAt

	return &todo, nil
}
