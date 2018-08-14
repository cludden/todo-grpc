// Package grpc provides a grpc server with supporint middleware
package grpc

import (
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
	log    logrus.FieldLogger
	port   uint32
	server *grpc.Server
}

// NewServer returns a new Server value
func NewServer(c *Config) (*Server, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// create server
	grpcServer := grpc.NewServer(
	// grpc.StreamInterceptor(middleware.ChainStreamServer(
	// 	grpcLogrus.StreamServerInterceptor(),
	// )),
	// grpc.UnaryInterceptor(middleware.ChainUnaryServer(
	// 	grpcLogrus.UnaryServerInterceptor(),
	// )),
	)
	proto.RegisterTodosServer(grpcServer, c.Server)
	s := Server{
		log:    c.Log,
		port:   c.Port,
		server: grpcServer,
	}
	return &s, nil
}

// Listen starts the underlying gRPC server instance and begins accepting
// connections on the configured port
func (s *Server) Listen() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		return fmt.Errorf("unable to create listener: %v", err)
	}
	s.log.Infof("grpc server listening on port %d", s.port)
	return s.server.Serve(lis)
}
