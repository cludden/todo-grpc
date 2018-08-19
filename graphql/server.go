// Package graphql provides the graphql server
package graphql

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"todo-grpc/graphql/generated"
	"todo-grpc/graphql/resolvers"
	"todo-grpc/proto"
	"todo-grpc/validation"

	"github.com/99designs/gqlgen/graphql"
	handler "github.com/99designs/gqlgen/handler"
	"github.com/sirupsen/logrus"
)

// Config describes the input to a NewServer operation
type Config struct {
	Graphiql bool
	Log      logrus.FieldLogger `validate:"required"`
	Port     uint32             `validate:"required"`
	Todos    proto.TodosClient  `validate:"required"`
}

// Server wraps the underlying graphql server and exposes methods
// for interacting with it
type Server struct {
	config *Config
	mux    *http.ServeMux
}

// NewServer returns a new server value
func NewServer(c *Config) (*Server, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// create root resolver
	resolver, err := resolvers.NewResolver(&resolvers.Config{
		Todos: c.Todos,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating graphql resolver: %v", err)
	}

	// create server
	s := Server{
		config: c,
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", s.GraphQL(resolver))
	if c.Graphiql {
		mux.HandleFunc("/", handler.Playground("todo", "/graphql"))
	}

	s.mux = mux
	return &s, nil
}

// GraphQL returns the graphql handler
func (s *Server) GraphQL(r *resolvers.Resolver) http.HandlerFunc {
	return handler.GraphQL(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: r,
		}),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetResolverContext(ctx)
			now := time.Now()
			fmt.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err, time.Since(now))
			return res, err
		}),
	)
}

// Listen starts the underlying graphql server
func (s *Server) Listen() error {
	s.config.Log.Infof("graphql server listening on port %d", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.mux)
}
