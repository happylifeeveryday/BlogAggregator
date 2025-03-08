-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name TEXT
);

-- +goose Down
DROP TABLE users;