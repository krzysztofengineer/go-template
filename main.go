package main

import (
	"log/slog"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
