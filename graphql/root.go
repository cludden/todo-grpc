package graphql

import (
	"todo-grpc/graphql/generated"
	"todo-grpc/proto"
)

// Root implements the root resolver interface
type Root struct {
	todos proto.TodosClient
}

// Mutations implements the root Mutation interface
type Mutations struct {
	todos proto.TodosClient
}

// Query implements the root Query interface
type Query struct {
	todos proto.TodosClient
}

// Mutation returns the root mutation resolver
func (r *Root) Mutation() generated.MutationResolver {
	resolver := Mutations{
		todos: r.todos,
	}
	return &resolver
}

// Query returns the root query resolver
func (r *Root) Query() generated.QueryResolver {
	return nil
}

// Todo returns the todo type resolver
func (r *Root) Todo() generated.TodoResolver {
	return nil
}
