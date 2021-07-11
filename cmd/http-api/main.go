package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/locnguyenvu/mangden/pkg/app"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/locnguyenvu/mangden/pkg/logger"
	"github.com/locnguyenvu/mangden/pkg/storage/mysql"
)

func main() {
	cfg, _ := config.New()

	l := logger.NewLogger(cfg.LogFormat, cfg.LogLevel)

	gormDB, _ := mysql.CreateOrm(cfg)

	deps := &app.Dependencies{
		Logger: l,
		GormDB: gormDB,
	}

	r := CreateRouter(deps)

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}

	go func(s *http.Server) {
		l.Infof("Server is running on: %s", cfg.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Fatalf("Failed to start server: %s", err)
		}
	}(server)

	stopCh := make(chan os.Signal, 2)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	l.Infof("Gracefully shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		l.WithError(err).Error("Error on shutting server down gracefully")
	}
}
