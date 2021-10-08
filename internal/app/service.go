package app

import (
	_ "github.com/lib/pq"
	"github.com/v1ack/linkShorter/internal/providers"
	desc "github.com/v1ack/linkShorter/pkg"
)

type Service interface {
	desc.ShorterServer
}

type service struct {
	provider providers.DataProvider
}

func NewService(provider providers.DataProvider) (Service, error) {
	return &service{provider: provider}, nil
}
