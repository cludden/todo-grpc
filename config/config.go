// Package config provides types and methods for loading environment configuration
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config describes the optional and required application configuration
type Config struct {
	Elasticsearch struct {
		Index string `validate:"required"`
		Type  string `validate:"required"`
		URL   string `validate:"required"`
	} `validate:"required"`
	GraphQL struct {
		Graphiql      bool
		Port          uint32 `validate:"required"`
		TodosEndpoint string `mapstructure:"todos-endpoint" validate:"required"`
	} `validate:"required"`
	GRPC struct {
		Port uint32 `validate:"required"`
	} `validate:"required"`
	GRPCGateway struct {
		Endpoint string `validate:"required"`
		Port     uint32 `validate:"required"`
	} `mapstructure:"grpc-gateway" validate:"required"`
	Log struct {
		Level string `validate:"required"`
	} `validate:"required"`
}

// New parses, loads, and validates application configuration from various
// environment sources
func New() (*Config, error) {
	var c Config

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/todos")
	viper.AddConfigPath(".")

	viper.SetDefault("elasticsearch.index", "todos")
	viper.SetDefault("elasticsearch.type", "doc")
	viper.SetDefault("graphql.port", 9000)
	viper.SetDefault("grpc.port", 10000)
	viper.SetDefault("grpc-gateway.port", 10001)
	viper.SetDefault("log.level", "info")

	viper.BindEnv("elasticsearch.index", "ELASTICSEARCH_INDEX")
	viper.BindEnv("elasticsearch.type", "ELASTICSEARCH_TYPE")
	viper.BindEnv("elasticsearch.url", "ELASTICSEARCH_URL")
	viper.BindEnv("graphql.graphiql", "GRAPHQL_GRAPHIQL")
	viper.BindEnv("graphql.port", "GRAPHQL_PORT")
	viper.BindEnv("graphql.todos-endpoint", "GRAPHQL_TODOS_ENDPOINT")
	viper.BindEnv("grpc.port", "GRPC_PORT")
	viper.BindEnv("grpc-gateway.endpoint", "GRPC_GATEWAY_ENDPOINT")
	viper.BindEnv("grpc-gateway.port", "GRPC_GATEWAY_PORT")
	viper.BindEnv("local.path", "LOCAL_PATH")
	viper.BindEnv("log.level", "LOG_LEVEL")

	viper.ReadInConfig()

	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %v", err)
	}

	return &c, nil
}
