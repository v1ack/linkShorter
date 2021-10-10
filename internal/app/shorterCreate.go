package app

import (
	"context"
	"errors"
	"github.com/v1ack/linkShorter/internal/store"
	desc "github.com/v1ack/linkShorter/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Create(ctx context.Context, request *desc.CreateLinkRequest) (*desc.CreateLinkResponse, error) {
	originalLink := request.Link
	if originalLink == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	hashiableLink := originalLink
	var shortLink string

	// Check for collision in loop
	// If collision exists - add "0" to hashiableLink to get a new hash
	for {
		shortLink = hashString(hashiableLink)

		err := s.provider.Create(ctx, originalLink, shortLink)

		// ShortLink created
		if err == nil {
			break
		}
		if errors.Is(err, store.ErrUnavailable) {
			return nil, status.Error(codes.Unavailable, store.ErrUnavailable.Error())
		}

		// Any error but already exists
		if !errors.Is(err, store.ErrAlreadyExists) {
			return nil, err
		}

		// Get exists link
		existsLink, err := s.provider.Get(ctx, shortLink)
		if err != nil {
			if errors.Is(err, store.ErrUnavailable) {
				return nil, status.Error(codes.Unavailable, store.ErrUnavailable.Error())
			}
			return nil, err
		}

		// If exists is the same as originalLink - success
		if existsLink == originalLink {
			break
		}

		hashiableLink += "0"
	}

	return &desc.CreateLinkResponse{
		ShortLink: shortLink,
	}, nil
}
