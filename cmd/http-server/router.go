package main

import (
    "net/http"

    "go.uber.org/dig"
	"github.com/go-chi/chi/v5"
    "mck.co/fuel/internal/testserver"
)

type Handlers struct {
	dig.In

	User *testserver.Handler
}

func loadRouter(handlers Handlers) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
    r.Get("/migrate", handlers.User.Migrate)
    r.Post("/users", handlers.User.CreateUser)
    r.Get("/users", handlers.User.ListUser)
	return r
}
