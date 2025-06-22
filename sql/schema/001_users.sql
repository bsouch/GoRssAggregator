-- +goose Up

CREATE TABLE Users (
    id UUID NOT NULL PRIMARY KEY UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down

DROP TABLE Users;