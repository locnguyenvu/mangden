package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(logFormat, defaultLevel string) *logrus.Logger {
	l := logrus.New()
	setFormat(l, logFormat)
	setLevel(l, defaultLevel)
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
