package main

import (
	"log/slog"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	guestMiddleware := NewMiddlewareStack(NewLoggerMiddleware(), NewRecoverMiddleware())

	guest := http.NewServeMux()
	guest.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})
	guest.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	r.Handle("/", guestMiddleware(guest))

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
