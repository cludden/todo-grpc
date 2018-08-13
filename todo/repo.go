package todo

import (
	"context"
	"todo-grpc/proto"
)

// Repository describes a value that provides access to a todo store
type Repository interface {
	CreateTodo(ctx context.Context, input *Todo) error
	ListTodos(ctx context.Context, input *proto.ListTodosInput) (*ListTodosOutput, error)
}

// ListTodosOutput describes the output from a successful ListTodos operation
type ListTodosOutput struct {
	Todos []*Todo
	Total int64
}
