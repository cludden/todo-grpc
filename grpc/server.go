// Package grpc provides a grpc server with supporting middleware
package grpc

import (
	"errors"
	"fmt"
	"net"
	"todo-grpc/proto"
	"todo-grpc/validation"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Config defines the input to a NewServer operation
type Config struct {
	Log    logrus.FieldLogger `validate:"required"`
	Port   uint32             `validate:"required"`
	Server proto.TodosServer  `validate:"required"`
}

// Server encapsulates an underlying gRPC server and provides methods for
// interacting with it
type Server struct {
	entry  *logrus.Entry
	port   uint32
	server *grpc.Server
}

// NewServer returns a new Server value
func NewServer(c *Config) (*Server, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	log, ok := c.Log.(*logrus.Entry)
	if !ok {
		return nil, errors.New("invalid logger provided")
	}

	// create server base server
	s := Server{
		entry: log,
		port:  c.Port,
	}

	// create grpc server
	grpcServer := grpc.NewServer(s.Interceptors()...)
	proto.RegisterTodosServer(grpcServer, c.Server)
	s.server = grpcServer
	return &s, nil
}

// Listen starts the underlying gRPC server instance and begins accepting
// connections on the configured port
func (s *Server) Listen() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		return fmt.Errorf("unable to create listener: %v", err)
	}
	s.entry.Infof("grpc server listening on port %d", s.port)
	return s.server.Serve(lis)
}
