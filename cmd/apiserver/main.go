package main


import (
    "context"
    "fmt"
    "os"
    "net/http"
    "os/signal"
    "syscall"

    "github.com/sirupsen/logrus"
    "go.uber.org/dig"
    "mck.co/fuel/pkg/database/mysql"
    "mck.co/fuel/pkg/logging"
    "mck.co/fuel/internal/routes"
)

var (
    serverAddress = "0.0.0.0:8000"
)

func main() {
    c := dig.New()

    if val, exists := os.LookupEnv("ADDR"); exists {
        serverAddress = val
    }

    constructors := []interface{}{
        logging.NewFromEnv,
        mysql.NewGormFromEnv,

        routes.APIServer,
    }

    for _, constructor := range constructors {
        if err := c.Provide(constructor); err != nil {
            fmt.Println("Failed to bootstrap", err.Error())
            os.Exit(1)
        }
    }

    c.Invoke(func(handler http.Handler, logger logrus.FieldLogger) {
        webserver := &http.Server{
            Addr: serverAddress,
            Handler: handler,
        }
        go func() {
            logger.Infof("Server is running on: %s", serverAddress)
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

