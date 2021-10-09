package store

import (
	"context"
)

type DataStore interface {
	Get(ctx context.Context, shortLink string) (string, error)
	Create(ctx context.Context, originalLink, shortLink string) error
}
