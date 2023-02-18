package logging

import (
    "context"

    "github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
    "github.com/sirupsen/logrus"

)

const (
    defaultFormat = "json"
    defaultLevel = "info"
)

type Config struct {
    LogFormat string `config:"LOG_FORMAT" validate:"oneof=text json"`
    LogLevel string `config:"LOG_LEVEL" validate:"oneof=debug info warn error fatal panic"`
}


func NewFromEnv() logrus.FieldLogger {
    cfg := &Config{
        defaultFormat,
        defaultLevel,
    }
	ctx := context.Background()
	loader := confita.NewLoader(env.NewBackend())
	loader.Load(ctx, cfg)
	l := logrus.New()
	defaultLogLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		l.SetLevel(defaultLogLevel)
	}
	if cfg.LogFormat == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}
	return l
}

