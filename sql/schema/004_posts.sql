-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title text NOT NULL,
    url text NOT null,
    description text NOT null,
    published_at text,
    feed_id uuid NOT NULL
);
-- +goose Down
DROP TABLE posts