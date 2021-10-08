-- name: GetLink :one
SELECT *
FROM links
WHERE short_link = $1
LIMIT 1;

-- name: CreateLink :one
INSERT INTO links (short_link, original_link)
VALUES ($1, $2)
RETURNING *;
