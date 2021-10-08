package providers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/v1ack/linkShorter/internal"
	"github.com/v1ack/linkShorter/internal/db"
)

type pgProvider struct {
	query      *db.Queries
	connection *sql.DB
}

func CreatePgProvider(connectionString string) (*pgProvider, error) {
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	query := db.New(connection)
	return &pgProvider{query: query, connection: connection}, nil
}

func (p *pgProvider) Get(ctx context.Context, shortLink string) (string, error) {
	link, err := p.query.GetLink(ctx, shortLink)
	if err == nil {
		return link.OriginalLink, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return "", internal.ErrNotFound
	}
	return "", internal.ErrUnavailable
}

func (p *pgProvider) Create(ctx context.Context, originalLink, shortLink string) error {
	_, err := p.query.CreateLink(ctx, db.CreateLinkParams{
		ShortLink:    shortLink,
		OriginalLink: originalLink,
	})

	return err
}

func (p *pgProvider) Close() {
	err := p.connection.Close()
	if err != nil {
		panic("Cant close connection")
	}
}
