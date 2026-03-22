-- +goose Up
CREATE TABLE if NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    create_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;