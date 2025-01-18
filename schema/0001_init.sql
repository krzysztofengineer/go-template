-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255)
    );

CREATE UNIQUE INDEX users_email ON users (email);

INSERT INTO
    users (email)
VALUES
    ("test@example.com");

-- +goose StatementEnd