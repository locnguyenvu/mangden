package main


import (
    "context"
    "fmt"
    "os"
    "net/http"
    "os/signal"
    "syscall"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/sirupsen/logrus"
    "go.uber.org/dig"
    "mck.co/fuel/pkg/db/mysql"
    "mck.co/fuel/internal/testserver"
)

func main() {
    c := dig.New()

	constructors := []interface{}{
        loadEnvConfig,
        initLogger,
        loadRouter,
        mysql.NewGorm,
        testserver.NewHandler,
	}

	for _, constructor := range constructors {
		if err := c.Provide(constructor); err != nil {
			fmt.Println("Failed to bootstrap", err.Error())
            os.Exit(1)
		}
	}

	c.Invoke(func(envConfig *EnvConfig, handler http.Handler, logger logrus.FieldLogger) {
        webserver := &http.Server{
            Addr: envConfig.ServerAddr,
            Handler: handler,
        }
        go func() {
            logger.Infof("Server is running on: %s", envConfig.ServerAddr)
            err := webserver.ListenAndServe()
            if err != nil {
                logger.Fatalf("%s", err)
            }
        }()
        stopCh := make(chan os.Signal, 2)
        signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
        <-stopCh
        logger.Infof("Gracefully shutting down server")
        if err := webserver.Shutdown(context.Background()); err != nil {
            logger.WithError(err).Error("Error on shutting server down gracefully")
        }
    })
}

type EnvConfig struct {
    ServerAddr string   `config:"ADDR"`
    LogFormat string    `config:"LOG_FORMAT" validate:"oneof=text json"`
    LogLevel string     `config:"LOG_LEVEL" validate:"oneof=debug info warn error fatal panic"`
}

func loadEnvConfig() *EnvConfig {
    defaultAddr := "0.0.0.0:8000"
    defaultLogLevel := "info"
    defaultLogFormat := "json"
    
	cfg := &EnvConfig{
		ServerAddr: defaultAddr,
		LogLevel: defaultLogLevel,
		LogFormat: defaultLogFormat,
	}

	ctx := context.Background()
	loader := confita.NewLoader(env.NewBackend())
	err := loader.Load(ctx, cfg)

	if err != nil {
		return cfg
	}

	return cfg
}

func initLogger(envConfig *EnvConfig) logrus.FieldLogger {
	var lg logrus.FieldLogger
	l := logrus.New()
	defaultLogLevel, err := logrus.ParseLevel(envConfig.LogLevel)
	if err != nil {
		l.SetLevel(defaultLogLevel)
	}
	if envConfig.LogFormat == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}
	lg = l
	return lg
}
