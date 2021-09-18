package logger

import (
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
)

func New(cfg *config.Config) logrus.FieldLogger {
	var lg logrus.FieldLogger
	l := logrus.New()
	defaultLogLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		l.SetLevel(defaultLogLevel)
	}
	if cfg.LogFormat == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}
	lg = l
	return lg
}
