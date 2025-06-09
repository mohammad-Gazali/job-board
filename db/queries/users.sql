-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (email, username, password)
VALUES ($1, $2, $3);