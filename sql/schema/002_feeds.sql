-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id uuid NOT NULL references users(id) ON DELETE CASCADE
    
);
ALTER TABLE feeds 
ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose Down
DROP TABLE feeds;