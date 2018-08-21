// Package gateway provides a configured gRPC gateway server that proxies json over
// http to protocol buffers over http2 (gRPC)
package gateway

import (
	"context"
	"fmt"
	"net/http"
	"todo-grpc/proto"
	"todo-grpc/validation"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Config describes the input to a NewServer operation
type Config struct {
	Endpoint string             `validate:"required"`
	Log      logrus.FieldLogger `validate:"required"`
	Port     uint32             `validate:"required"`
}

// Server encapsulates an underlying grpc gateway server and exposes methods
// for interacting with it
type Server struct {
	cancel context.CancelFunc
	log    logrus.FieldLogger
	mux    *runtime.ServeMux
	port   uint32
}

// NewServer returns a new server value
func NewServer(c *Config) (*Server, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// create root context
	ctx, cancel := context.WithCancel(context.Background())

	// register handlers
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterTodosHandlerFromEndpoint(ctx, mux, c.Endpoint, opts)
	if err != nil {
		defer cancel()
		return nil, fmt.Errorf("unable to register gateway handler: %v", err)
	}

	s := Server{
		cancel: cancel,
		log:    c.Log,
		mux:    mux,
		port:   c.Port,
	}
	return &s, nil
}

// Listen starts the underlying gateway server and begins accepting connections
// on the configured port
func (s *Server) Listen() error {
	defer s.cancel()
	s.log.Infof("grpc-gateway server listening on port %d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}
