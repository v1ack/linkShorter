package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/v1ack/linkShorter/internal/db"
	"strings"
)

type dbStore struct {
	query      *db.Queries
	connection *sql.DB
}

func CreateDbProvider(connection *sql.DB) *dbStore {
	query := db.New(connection)
	return &dbStore{query: query, connection: connection}
}

func (p *dbStore) Get(ctx context.Context, shortLink string) (string, error) {
	link, err := p.query.GetLink(ctx, shortLink)
	if err == nil {
		return link.OriginalLink, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrNotFound
	}
	return "", ErrUnavailable
}

func (p *dbStore) Create(ctx context.Context, originalLink, shortLink string) error {
	_, err := p.query.CreateLink(ctx, db.CreateLinkParams{
		ShortLink:    shortLink,
		OriginalLink: originalLink,
	})
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return ErrAlreadyExists
	}
	return err
}
