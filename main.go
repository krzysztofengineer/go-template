package main

import (
	"log/slog"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	guestMiddleware := NewMiddlewareStack(NewLoggerMiddleware())

	guest := http.NewServeMux()
	guest.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	r.Handle("/", guestMiddleware(guest))

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
