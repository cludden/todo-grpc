package todo

import (
	"context"
	"fmt"
	"time"
	"todo-grpc/proto"
	"todo-grpc/validation"

	"gopkg.in/guregu/null.v3"
)

// Server implements the mindflash.todos.Todos service interface
type Server struct {
	repo Repository
}

// NewServer returns a new server value
func NewServer(repo Repository) (*Server, error) {
	s := Server{
		repo: repo,
	}
	return &s, nil
}

// CreateTodo creates a new todo task
func (s *Server) CreateTodo(ctx context.Context, in *proto.CreateTodoInput) (*proto.CreateTodoOutput, error) {
	// validate input
	if err := validation.Validate.Struct(in); err != nil {
		return nil, fmt.Errorf("invalid input: %v", err)
	}

	// create todo
	now := time.Now()
	record := Todo{
		ID:          now.Format(time.RFC3339Nano),
		Complete:    false,
		CompletedAt: null.NewTime(time.Time{}, false),
		CreatedAt:   now,
		Description: null.StringFrom(in.GetDescription()),
		Title:       in.Title,
	}
	if err := s.repo.CreateTodo(ctx, &record); err != nil {
		return nil, fmt.Errorf("unexpected error creating todo: %v", err)
	}

	// marshal output
	todo, err := record.MarshalPB()
	if err != nil {
		return nil, fmt.Errorf("error marshalling output: %v", err)
	}

	// define output
	out := proto.CreateTodoOutput{
		Todo: todo,
	}
	return &out, nil
}

// ListTodos retrieves a paginated list of todos
func (s *Server) ListTodos(ctx context.Context, in *proto.ListTodosInput) (*proto.ListTodosOutput, error) {
	// validate input
	if err := validation.Validate.Struct(in); err != nil {
		return nil, fmt.Errorf("invalid input: %v", err)
	}

	// retrieve todos
	res, err := s.repo.ListTodos(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("unexpected error retrieving todos: %v", err)
	}

	// define output
	out := proto.ListTodosOutput{
		Todos: make([]*proto.Todo, res.Total),
		Total: res.Total,
	}

	// marshal outout
	for i, t := range res.Todos {
		todo, err := t.MarshalPB()
		if err != nil {
			return nil, fmt.Errorf("error marshalling output: %v", err)
		}
		out.Todos[i] = todo
	}

	return &out, err
}
