package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locnguyenvu/mangden/internal/user"
	"go.uber.org/dig"
)

type HandlerParams struct {
	dig.In

	UserHandler *user.HttpHandler
}

type responseWrapper func(w http.ResponseWriter, r *http.Request) error

func (rw responseWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := rw(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewRouter(p HandlerParams) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mangden project"))
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error page", 500)
		return
	})

	r.Get("/user", p.UserHandler.CreateUser)

	return r
}
