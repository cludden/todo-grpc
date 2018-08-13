// Package config provides types and methods for loading environment configuration
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config describes the optional and required application configuration
type Config struct {
	GRPC struct {
		Port uint32 `validate:"required"`
	} `validate:"required"`
	GRPCGateway struct {
		Endpoint string `validate:"required,min=3"`
		Port     uint32 `validate:"required"`
	} `mapstructure:"grpc-gateway" validate:"required"`
	Local struct {
		Path string `validate:"required"`
	} `validate:"required"`
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

	viper.SetDefault("grpc.port", 10000)
	viper.SetDefault("grpc-gateway.port", 10001)
	viper.SetDefault("local.path", "todos.db")
	viper.SetDefault("log.level", "info")

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