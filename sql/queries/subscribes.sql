-- name: CreateSubscribe :one
INSERT INTO subscribes (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSubscribeByUserID :many
SELECT * FROM subscribes WHERE user_id = $1;

-- name: DeleteSubscribe :one
DELETE FROM subscribes WHERE id = $1 AND user_id = $2
RETURNING *;