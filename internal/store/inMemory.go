package store

import (
	"context"
	"sync"
)

type inMemoryStore struct {
	mu    *sync.Mutex
	store map[string]string
}

func (p *inMemoryStore) Get(_ context.Context, shortLink string) (string, error) {
	link, ok := p.store[shortLink]
	if ok == false {
		return "", ErrNotFound
	}
	return link, nil
}

func (p *inMemoryStore) Create(_ context.Context, originalLink, shortLink string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, exists := p.store[shortLink]
	if exists {
		return ErrAlreadyExists
	}

	p.store[shortLink] = originalLink
	return nil
}

func CreateInMemoryProvider() *inMemoryStore {
	return &inMemoryStore{store: map[string]string{}, mu: &sync.Mutex{}}
}
