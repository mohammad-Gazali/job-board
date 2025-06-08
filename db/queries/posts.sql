-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GetPostById :one
SELECT * FROM posts WHERE id = $1;

-- name: CreatePost :exec
INSERT INTO posts (title)
VALUES ($1);