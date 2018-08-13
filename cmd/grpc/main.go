// GRPC server entrypoint
package main

import (
	"log"
	"os"
	"todo-grpc/config"
	"todo-grpc/resolver"
)

func main() {
	// load configuration
	c, err := config.New()
	handleError(err)

	// create resolver
	r := resolver.NewResolver(c)
	server, err := r.GRPCServer()
	handleError(err)
	handleError(server.Listen())
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
