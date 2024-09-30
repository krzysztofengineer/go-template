package main

import (
	"database/sql"
	"io/fs"
	"log/slog"
	"slices"
	"sort"
)

func MigrateDB(db *sql.DB) ([]string, error) {
	applied := []string{}
	// create migrations table if doesn't exist
	sql := `
	CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(sql); err != nil {
		return applied, err
	}

	files, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return applied, err
	}

	names := make([]string, len(files))
	for i, file := range files {
		names[i] = file.Name()
	}
	sort.Strings(names)

	alreadyApplied, err := appliedMigrations(db)
	if err != nil {
		return applied, err
	}

	for _, name := range names {
		if slices.Contains(alreadyApplied, name) {
			continue
		}

		content, err := fs.ReadFile(migrationsFS, "migrations/"+name)
		if err != nil {
			return applied, err
		}

		if _, err := db.Exec(string(content)); err != nil {
			slog.Error("failed to apply migration", "name", name, "error", err)
			return applied, err
		}

		if _, err := db.Exec("INSERT INTO migrations (name) VALUES (?)", name); err != nil {
			slog.Error("failed to record migration", "name", name, "error", err)
			return applied, err
		}

		slog.Info("\033[32mOK\033[0m", "name", name)
		applied = append(applied, name)
	}

	return applied, nil
}

func appliedMigrations(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	names := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}
