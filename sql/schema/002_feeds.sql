-- +goose UP
CREATE TABLE feeds (
    id INTEGER
    created_at TIMESTAMP
    updated_at TIMESTAMP
    name TEXT,
    url UNIQUE TEXT,
    user_id uuid references users ON DELETE CASCADE
    

) -- +goose Down
DROP TABLE feeds