package todo

import (
	"context"
	"fmt"
	"time"
	"todo-grpc/proto"
	"todo-grpc/validation"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/guregu/null.v3"
)

// Config describes the input to a NewServer operation
type Config struct {
	Repository Repository `validate:"required"`
}

// Server implements the mindflash.todos.Todos service interface
type Server struct {
	repo Repository
}

// NewServer returns a new server value
func NewServer(c *Config) (*Server, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	s := Server{
		repo: c.Repository,
	}
	return &s, nil
}

// CompleteTodo marks an existing todo as completed
func (s *Server) CompleteTodo(ctx context.Context, in *proto.CompleteTodoInput) (*proto.CompleteTodoOutput, error) {
	// validate input
	if err := validation.Validate.Struct(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid input: %v", err)
	}

	// update todo
	todo, err := s.repo.CompleteTodo(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	// marshal todo
	marshalled, err := todo.MarshalPB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error marshalling output: %v", err)
	}

	// define output
	out := proto.CompleteTodoOutput{
		Todo: marshalled,
	}
	return &out, nil
}

// CreateTodo creates a new todo task
func (s *Server) CreateTodo(ctx context.Context, in *proto.CreateTodoInput) (*proto.CreateTodoOutput, error) {
	// validate input
	if err := validation.Validate.Struct(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid input: %v", err)
	}

	// create todo
	record := Todo{
		ID:          uuid.NewV4().String(),
		Complete:    false,
		CompletedAt: null.NewTime(time.Time{}, false),
		CreatedAt:   time.Now(),
		Description: null.StringFrom(in.GetDescription()),
		Title:       in.Title,
	}
	if err := s.repo.CreateTodo(ctx, &record); err != nil {
		return nil, err
	}

	// marshal output
	todo, err := record.MarshalPB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error marshalling output: %v", err)
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
		return nil, status.Errorf(codes.InvalidArgument, "invalid input: %v", err)
	} else if in.First == 0 {
		in.First = 20
	}

	// retrieve todos
	res, err := s.repo.ListTodos(ctx, in)
	if err != nil {
		return nil, err
	}

	// define output
	out := proto.ListTodosOutput{
		Todos: make([]*proto.Todo, len(res.Todos)),
		Total: res.Total,
	}

	// marshal outout
	for i, t := range res.Todos {
		todo, err := t.MarshalPB()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error marshalling output: %v", err)
		}
		out.Todos[i] = todo
	}

	return &out, err
}
