package grpc

// import (
// 	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
// 	"github.com/sirupsen/logrus"
// )

// LogRequests returns
// func (s *Server) LogRequests(unary ) {
// 	// create logrus entry
// 	log, ok := s.log.(*logrus.Entry)
// 	if !ok {
// 		log = logrus.WithField("pkg", "grpc")
// 	}
// 	entry := logrus.NewEntry(log.Logger)

// 	// create logging options slice
// 	options := []grpcLogrus.Option{
// 		grpcLogrus.WithLevels(grpcLogrus.DefaultCodeToLevel),
// 	}
// }
