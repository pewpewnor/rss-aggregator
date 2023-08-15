-- name: CreateSubscribe :one
INSERT INTO subscribes (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
