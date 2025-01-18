package main

import (
	"go-template/pages"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			pages.Home().Render(r.Context(), w)
		})
	})

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
