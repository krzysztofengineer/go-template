package main

import (
	"database/sql"
	"embed"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed views
	viewsFS embed.FS

	//go:embed static
	staticFS embed.FS

	//go:embed migrations/*.sql
	migrationsFS embed.FS
)

func main() {
	if _, err := os.Stat("db.sqlite"); os.IsNotExist(err) {
		f, err := os.Create("db.sqlite")
		if err != nil {
			slog.Error("failed to create database", "error", err)
			os.Exit(1)
		}
		f.Close()
	}

	db, err := sql.Open("sqlite3", "file:db.sqlite")
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	applied, err := MigrateDB(db)
	if err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	if len(applied) > 0 {
		slog.Info("applied migrations", "migrations", applied)
	} else {
		slog.Info("database is up to date")
	}

	r := http.NewServeMux()

	guestMiddleware := NewMiddlewareStack(NewLoggerMiddleware(), NewRecoverMiddleware())
	guest := http.NewServeMux()

	homeHandler := NewHomeHandler()
	guest.HandleFunc("GET /{$}", homeHandler.Home())

	r.Handle("/", guestMiddleware(guest))
	r.Handle("/static/", http.FileServer(http.FS(staticFS)))

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
