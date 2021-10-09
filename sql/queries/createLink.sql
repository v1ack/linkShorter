-- name: CreateLink :one
INSERT INTO links (short_link, original_link)
VALUES ($1, $2)
RETURNING *;
