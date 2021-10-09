-- name: GetLink :one
SELECT *
FROM links
WHERE short_link = $1
LIMIT 1;