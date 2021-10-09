package app

import (
	"github.com/v1ack/linkShorter/internal/store"
	desc "github.com/v1ack/linkShorter/pkg"
)

type Service interface {
	desc.ShorterServer
}

type service struct {
	provider store.DataStore
}

func NewService(provider store.DataStore) (Service, error) {
	return &service{provider: provider}, nil
}
