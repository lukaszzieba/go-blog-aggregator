-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name FROM users WHERE name = $1;

-- name: GetUsers :many
SELECT id, created_at, updated_at, name FROM users;

-- name: DeleteAllUsers :exec
DELETE FROM USERS;
