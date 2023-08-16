-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, url, title, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetNewestPostsForUser :many
SELECT posts.*
FROM posts JOIN subscribes ON posts.feed_id = subscribes.feed_id
WHERE subscribes.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;