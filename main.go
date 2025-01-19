package main

import (
	"flag"
	"fmt"
	"go-template/database"
	"go-template/handler"
	"go-template/static"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	dsn  = flag.String("dsn", ":memory:", "sqlite dsn")
	port = flag.Int("port", 4000, "http server port")
)

func main() {
	flag.Parse()

	db := database.New(*dsn)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(noCacheMiddleware)

	homeHandler := handler.NewHome(db)

	r.Group(func(r chi.Router) {
		r.Get("/", homeHandler.Index)
		r.Post("/increase", homeHandler.Increase)
		r.Post("/decrease", homeHandler.Decrease)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(static.FS)))

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: r,
	}

	slog.Info(fmt.Sprintf("http://localhost:%d", *port))

	log.Fatal(s.ListenAndServe())
}
