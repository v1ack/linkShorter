// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const createLink = `-- name: CreateLink :one
INSERT INTO links (short_link, original_link)
VALUES ($1, $2)
RETURNING id, short_link, original_link
`

type CreateLinkParams struct {
	ShortLink    string
	OriginalLink string
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.db.QueryRowContext(ctx, createLink, arg.ShortLink, arg.OriginalLink)
	var i Link
	err := row.Scan(&i.ID, &i.ShortLink, &i.OriginalLink)
	return i, err
}

const getLink = `-- name: GetLink :one
SELECT id, short_link, original_link
FROM links
WHERE short_link = $1
LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, shortLink string) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLink, shortLink)
	var i Link
	err := row.Scan(&i.ID, &i.ShortLink, &i.OriginalLink)
	return i, err
}
