package logger

import (
	"time"

	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
)

func New(config *config.Config) *logrus.Logger {

	l := logrus.New()
	setFormat(l, config.LogFormat)
	setLevel(l, config.LogLevel)
	return l
}

func setFormat(l *logrus.Logger, format string) {
	m := logrus.FieldMap{
		logrus.FieldKeyTime: "@timestamp",
		logrus.FieldKeyMsg:  "message",
	}
	switch format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{
			FieldMap:        m,
			TimestampFormat: time.RFC3339Nano,
		})
	default:
		l.SetFormatter(&logrus.TextFormatter{
			FieldMap:        m,
			TimestampFormat: time.RFC3339Nano,
		})
	}
}

func setLevel(l *logrus.Logger, level string) {
	switch level {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}
}
