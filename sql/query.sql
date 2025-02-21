-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one 
INSERT INTO authors (email, password)
VALUES ($1, $2)
RETURNING id, email;


