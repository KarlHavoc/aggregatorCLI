-- name: CreateFeedFollows :one
WITH inserted_feed_follows AS (
    INSERT INTO
        feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES
        ($1, $2, $3, $4, $5) RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM
    inserted_feed_follows
    INNER JOIN users ON users.id = user_id
    INNER JOIN feeds ON feeds.id = feed_id;
-- name: GetFeedFollowsForUser :many
SELECT
    *
FROM
    feed_follows
WHERE
    user_id = $1;
-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 and user_id = $2;
