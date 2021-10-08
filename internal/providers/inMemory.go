package providers

import (
	"context"
	"github.com/v1ack/linkShorter/internal"
)

type inMemoryProvider struct {
	store map[string]string
}

func (p *inMemoryProvider) Close() {
	return
}

func (p *inMemoryProvider) Get(_ context.Context, shortLink string) (string, error) {
	link, ok := p.store[shortLink]
	if ok == false {
		return "", internal.ErrNotFound
	}
	return link, nil
}

func (p *inMemoryProvider) Create(_ context.Context, originalLink, shortLink string) error {
	_, exists := p.store[shortLink]
	if exists {
		return internal.ErrAlreadyExists
	}

	p.store[shortLink] = originalLink
	return nil
}

func CreateInMemoryProvider() *inMemoryProvider {
	return &inMemoryProvider{store: map[string]string{}}
}
