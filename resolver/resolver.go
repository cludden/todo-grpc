// Package resolver provides types and methods for loading application components
package resolver

import (
	"fmt"
	"todo-grpc/config"
	"todo-grpc/elasticsearch"
	gw "todo-grpc/gateway"
	"todo-grpc/graphql"
	"todo-grpc/grpc"
	"todo-grpc/proto"
	"todo-grpc/todo"

	"github.com/sirupsen/logrus"
	g "google.golang.org/grpc"
)

var version string

// Resolver provides methods for loading lazily loading application components
type Resolver struct {
	config                  config.Config
	elasticsearchRepository todo.Repository
	log                     logrus.FieldLogger
	graphqlServer           *graphql.Server
	grpcServer              *grpc.Server
	grpcGatewayServer       *gw.Server
	todosClient             proto.TodosClient
	todosServer             proto.TodosServer
}

// NewResolver returns a new resolver value
func NewResolver(c *config.Config) *Resolver {
	r := Resolver{
		config: *c,
	}
	return &r
}

// Config returns a copy of the application configuration
func (r *Resolver) Config() config.Config {
	return r.config
}

// ElasticsearchRepository returns a singleton local repository value
func (r *Resolver) ElasticsearchRepository() (todo.Repository, error) {
	if r.elasticsearchRepository == nil {
		c := r.Config()
		repo, err := elasticsearch.NewRepository(&elasticsearch.Config{
			Index:    c.Elasticsearch.Index,
			Log:      r.Log().WithField("pkg", "elasticsearch"),
			TypeName: c.Elasticsearch.Type,
			URL:      c.Elasticsearch.URL,
		})
		if err != nil {
			return nil, fmt.Errorf("error resolving elasticsearch repository: %v", err)
		}
		r.elasticsearchRepository = repo
	}
	return r.elasticsearchRepository, nil
}

// GraphQLServer returns a singleton graphql server value
func (r *Resolver) GraphQLServer() (*graphql.Server, error) {
	if r.graphqlServer == nil {
		c := r.Config()
		todos, err := r.TodosClient()
		if err != nil {
			return nil, err
		}
		server, err := graphql.NewServer(&graphql.Config{
			Graphiql: c.GraphQL.Graphiql,
			Log:      r.Log().WithField("pkg", "graphql"),
			Port:     c.GraphQL.Port,
			Todos:    todos,
		})
		if err != nil {
			return nil, fmt.Errorf("error resolving graphql server: %v", err)
		}
		r.graphqlServer = server
	}
	return r.graphqlServer, nil
}

// GRPCServer returns a singleton grpc server value
func (r *Resolver) GRPCServer() (*grpc.Server, error) {
	if r.grpcServer == nil {
		c := r.Config()
		s, err := r.TodosServer()
		if err != nil {
			return nil, err
		}
		server, err := grpc.NewServer(&grpc.Config{
			Log:    r.Log(),
			Port:   c.GRPC.Port,
			Server: s,
		})
		if err != nil {
			return nil, fmt.Errorf("error resolving grpc server: %v", err)
		}
		r.grpcServer = server
	}
	return r.grpcServer, nil
}

// GatewayServer returns a singleton grpc gateway server value
func (r *Resolver) GatewayServer() (*gw.Server, error) {
	if r.grpcGatewayServer == nil {
		c := r.Config()
		server, err := gw.NewServer(&gw.Config{
			Endpoint: c.GRPCGateway.Endpoint,
			Log:      r.Log(),
			Port:     c.GRPCGateway.Port,
		})
		if err != nil {
			return nil, fmt.Errorf("error resolving grpc-gateway server: %v", err)
		}
		r.grpcGatewayServer = server
	}
	return r.grpcGatewayServer, nil
}

// Log returns a singleton logger instance
func (r *Resolver) Log() logrus.FieldLogger {
	if r.log == nil {
		c := r.Config()
		logrus.SetFormatter(&logrus.JSONFormatter{})
		level, err := logrus.ParseLevel(c.Log.Level)
		if err != nil {
			level = logrus.InfoLevel
		}
		logrus.SetLevel(level)
		r.log = logrus.WithFields(logrus.Fields{
			"name":    "todo",
			"version": version,
		})
	}
	return r.log
}

// TodosClient returns a singleton todo client value
func (r *Resolver) TodosClient() (proto.TodosClient, error) {
	if r.todosClient == nil {
		c := r.Config()
		conn, err := g.Dial(c.GraphQL.TodosEndpoint, g.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("error resolving todos client: %v", err)
		}
		r.todosClient = proto.NewTodosClient(conn)
	}
	return r.todosClient, nil
}

// TodosServer returns a singleton todo server value
func (r *Resolver) TodosServer() (proto.TodosServer, error) {
	if r.todosServer == nil {
		repo, err := r.ElasticsearchRepository()
		if err != nil {
			return nil, err
		}
		server, err := todo.NewServer(&todo.Config{
			Repository: repo,
		})
		if err != nil {
			return nil, fmt.Errorf("error resolving todo server: %v", err)
		}
		r.todosServer = server
	}
	return r.todosServer, nil
}
