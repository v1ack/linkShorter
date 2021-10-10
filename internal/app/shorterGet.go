package app

import (
	"context"
	"errors"
	"github.com/v1ack/linkShorter/internal/store"
	desc "github.com/v1ack/linkShorter/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Get(ctx context.Context, request *desc.GetLinkRequest) (*desc.GetLinkResponse, error) {
	if request.Link == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	originalLink, err := s.provider.Get(ctx, request.Link)

	if err == nil {
		return &desc.GetLinkResponse{OriginalLink: originalLink}, nil
	}
	if errors.Is(err, store.ErrUnavailable) {
		return nil, status.Error(codes.Unavailable, store.ErrUnavailable.Error())
	}
	if errors.Is(err, store.ErrNotFound) {
		return nil, status.Error(codes.NotFound, store.ErrNotFound.Error())
	}
	return nil, err
}
