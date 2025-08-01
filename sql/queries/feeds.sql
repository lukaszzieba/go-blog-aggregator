-- name: CreateFeed :one
INSERT INTO feeds (name, created_at, updated_at ,url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5

)
RETURNING *;

-- name: GetFeedsWithUsers :many
SELECT f.name, f.url, u.name as user_name
FROM feeds f
JOIN users u ON f.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT f.id, f.name, f.created_at, f.updated_at, f.url
FROM feeds f
WHERE f.url  = $1; 
