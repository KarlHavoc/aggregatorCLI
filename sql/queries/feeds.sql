-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: DeleteFeeds :exec
DELETE FROM feeds;
-- name: GetFeeds :many
SELECT * FROM feeds;
-- name: GetFeed :one
SELECT id FROM feeds
WHERE url = $1;
-- name: GetFeedName :one
SELECT name FROM feeds
WHERE id = $1;


