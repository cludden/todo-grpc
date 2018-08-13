package grpc

import (
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"google.golang.org/grpc"
)

// Interceptors returns the default grpc middleware chain
func (s *Server) Interceptors() []grpc.ServerOption {
	// create default chains
	unary := []grpc.UnaryServerInterceptor{}
	stream := []grpc.StreamServerInterceptor{}

	s.LogRequests(&unary, &stream)

	options := []grpc.ServerOption{
		middleware.WithStreamServerChain(stream...),
		middleware.WithUnaryServerChain(unary...),
	}
	return options
}

// LogRequests appends logging middleware to both interceptor chains
func (s *Server) LogRequests(unary *[]grpc.UnaryServerInterceptor, stream *[]grpc.StreamServerInterceptor) {
	// initialize log options
	options := []grpcLogrus.Option{
		grpcLogrus.WithLevels(grpcLogrus.DefaultCodeToLevel),
	}

	*unary = append(*unary, grpcLogrus.UnaryServerInterceptor(s.entry, options...))
	*stream = append(*stream, grpcLogrus.StreamServerInterceptor(s.entry, options...))
}
