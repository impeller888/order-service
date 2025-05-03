package config

import (
	"fmt"

	"local/order-service/pkg/postgres"

	"github.com/aranw/yamlcfg"
)

type (
	Config struct {
		HTTP    HTTP    `yaml:"http"`
		Log     Log     `yaml:"log"`
		DB      DB      `yaml:"db"`
		Metrics Metrics `yaml:"metrics"`
		Swagger Swagger `yaml:"swagger"`
		Tracing Tracing `yaml:"tracing"`
	}

	HTTP struct {
		Addr string `yaml:"address"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	DB struct {
		PG postgres.Config `yaml:"postgres"`
	}

	Metrics struct {
		Enabled bool `yaml:"enabled"`
	}

	Swagger struct {
		Enabled bool `yaml:"enabled"`
	}

	Tracing struct {
		Enabled bool `yaml:"enabled"`
	}
)

// NewConfig returns app config.
func NewConfig(name string) (*Config, error) {
	if name == "" {
		name = "config/config.yaml"
	}
	cfg, err := yamlcfg.Parse[Config](name)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
