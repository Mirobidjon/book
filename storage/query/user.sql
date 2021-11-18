-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetUsers :many
SELECT *
FROM users OFFSET $1
LIMIT $2;
-- name: CreateUser :one
INSERT INTO users (
        id,
        first_name,
        last_name,
        email
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateUser :exec
UPDATE users
SET first_name = $1,
    last_name = $2,
    email = $3
WHERE id = $4;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
