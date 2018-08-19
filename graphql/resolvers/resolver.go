package resolvers

import (
	"fmt"
	"todo-grpc/graphql/generated"
	"todo-grpc/proto"
	"todo-grpc/validation"
)

// Config describes the input to a NewConfig operation
type Config struct {
	Todos proto.TodosClient `validate:"required"`
}

// Resolver implements the root resolver interface
type Resolver struct {
	todos proto.TodosClient
}

// NewResolver returns a new graphql root resolver value
func NewResolver(c *Config) (*Resolver, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	r := Resolver{
		todos: c.Todos,
	}
	return &r, nil
}

// Mutations implements the root Mutation interface
type Mutations Resolver

// Query implements the root Query interface
type Query Resolver

// Mutation returns the root mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver {
	return (*Mutations)(r)
}

// Query returns the root query resolver
func (r *Resolver) Query() generated.QueryResolver {
	return (*Query)(r)
}

// Todo returns the todo type resolver
func (r *Resolver) Todo() generated.TodoResolver {
	return (*Todos)(r)
}
