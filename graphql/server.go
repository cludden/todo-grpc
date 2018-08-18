package graphql

import (
	"fmt"
	"net/http"
	"todo-grpc/validation"
)

// Config describes the input to a NewServer operation
type Config struct {
	Port uint32 `validate:"required"`
}

// Server wraps the underlying graphql server and exposes methods
// for interacting with it
type Server struct {
	mux  *http.ServeMux
	port uint32
}

// NewServer returns a new server value
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	s := Server{
		port: c.Port,
	}

	mux := http.NewServeMux()

	s.mux = mux
	return &s, nil
}
