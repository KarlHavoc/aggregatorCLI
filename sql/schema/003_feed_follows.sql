-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id uuid NOT NULL UNIQUE references users(id) ON DELETE CASCADE,
    feed_id uuid NOT NULL UNIQUE references feeds(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feed_follows;