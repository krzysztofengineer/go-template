package main

import (
	"embed"
	"log/slog"
	"net/http"
)

var (
	//go:embed views
	viewsFS embed.FS

	//go:embed static
	staticFS embed.FS
)

func main() {
	r := http.NewServeMux()

	guestMiddleware := NewMiddlewareStack(NewLoggerMiddleware(), NewRecoverMiddleware())
	guest := http.NewServeMux()

	homeHandler := NewHomeHandler()
	guest.HandleFunc("GET /{$}", homeHandler.Home())
	guest.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	r.Handle("/", guestMiddleware(guest))
	r.Handle("/static/", http.FileServer(http.FS(staticFS)))

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
