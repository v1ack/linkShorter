package providers

import (
	"context"
)

type DataProvider interface {
	Get(ctx context.Context, shortLink string) (string, error)
	Create(ctx context.Context, originalLink, shortLink string) error
	Close()
}
