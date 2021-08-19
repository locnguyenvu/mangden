package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
)

type WebApp struct {
	config *config.Config
	server *http.Server
	logger *logrus.Logger
}

func NewWebApp(config *config.Config, logger *logrus.Logger, handler http.Handler) *WebApp {
	server := &http.Server{
		Addr:    config.Addr,
		Handler: handler,
	}
	return &WebApp{config, server, logger}
}

func (a *WebApp) Run() {
	go func(a *WebApp) {
		a.logger.Infof("Server is running on: %s", a.config.Addr)
		err := a.server.ListenAndServe()
		if err != nil {
			a.logger.Fatalf("Failed to start server: %s", err)
		}
	}(a)
	stopCh := make(chan os.Signal, 2)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	a.logger.Infof("Gracefully shutting down server")
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.logger.WithError(err).Error("Error on shutting server down gracefully")
	}
}
