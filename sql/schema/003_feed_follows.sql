-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id uuid NOT NULL references users(id) ON DELETE CASCADE,
    feed_id uuid NOT NULL references feeds(id) ON DELETE CASCADE,
    CONSTRAINT uq_user_id_feed_id UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;