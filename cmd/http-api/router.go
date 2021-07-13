package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locnguyenvu/mangden/internal/appconfig"
	"github.com/locnguyenvu/mangden/pkg/app"
	"github.com/zalora/prometheus-client-go/monitoring"
)

func CreateRouter(deps *app.Dependencies) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mangden project"))
	})

	monitoring := monitoring.New()

	r.Handle("/metrics", monitoring.MetricsHandler())

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error page", 500)
		return
	})

	appConfigRepository := appconfig.NewRepository(deps.GormDB)

	r.Mount("/config", appconfig.NewHandler(
		appConfigRepository,
	))

	return r
}
