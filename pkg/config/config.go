package config

import (
	"context"
	"fmt"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

const (
	defaultLogLevel  = "info"
	defaultLogFormat = "json"
	defaultAddr      = "0.0.0.0:8000"
	defaultDbPort    = 3306
)

type Config struct {
	Addr       string `config:"ADDR"`
	LogLevel   string `config:"LOG_LEVEL" validate:"oneof=debug info warn error fatal panic"`
	LogFormat  string `config:"LOG_FORMAT" validate:"oneof=text json"`
	DbHost     string `config:"DB_HOST"`
	DbUser     string `config:"DB_USER"`
	DbPassword string `config:"DB_PASSWORD"`
	DbPort     int    `config:"DB_PORT"`
	DbName     string `config:"DB_NAME"`
}

func New() (*Config, error) {
	cfg := &Config{
		Addr:      defaultAddr,
		LogLevel:  defaultLogLevel,
		LogFormat: defaultLogFormat,
		DbPort:    defaultDbPort,
	}

	ctx := context.Background()
	loader := confita.NewLoader(env.NewBackend())
	err := loader.Load(ctx, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) DbConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}
