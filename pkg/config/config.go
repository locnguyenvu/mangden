package config

import (
	"context"
	"fmt"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

const (
	defaultAddr                 = "0.0.0.0:8000"
	defaultEnvironment          = "local"
	defaultLogLevel             = "info"
	defaultLogFormat            = "json"
	defaultDbPort               = 3306
	defaultDBMaxIdleConnections = 10
	defaultDBMaxOpenConnections = 5
	defaultDBMaxConnLifeTime    = 30 * time.Minute
)

type Config struct {
	Addr                 string        `config:"ADDR"`
	Environment          string        `config:"ENVIRONMENT"`
	LogLevel             string        `config:"LOG_LEVEL" validate:"oneof=debug info warn error fatal panic"`
	LogFormat            string        `config:"LOG_FORMAT" validate:"oneof=text json"`
	DbHost               string        `config:"DB_HOST"`
	DbUser               string        `config:"DB_USER"`
	DbPassword           string        `config:"DB_PASSWORD"`
	DbPort               int           `config:"DB_PORT"`
	DbName               string        `config:"DB_NAME"`
	DBMaxIdleConnections int           `config:"DB_MAX_IDLE_CONNECTIONS"`
	DBMaxOpenConnections int           `config:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxConnLifetime    time.Duration `config:"DB_MAX_CONN_LIFETIME"`
}

func New() (*Config, error) {
	cfg := &Config{
		Addr:                 defaultAddr,
		Environment:          defaultEnvironment,
		LogLevel:             defaultLogLevel,
		LogFormat:            defaultLogFormat,
		DbPort:               defaultDbPort,
		DBMaxIdleConnections: defaultDBMaxIdleConnections,
		DBMaxOpenConnections: defaultDBMaxOpenConnections,
		DBMaxConnLifetime:    defaultDBMaxConnLifeTime,
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
