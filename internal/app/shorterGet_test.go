package app

import (
	"context"
	"fmt"
	"github.com/v1ack/linkShorter/internal/store"
	linkShorter "github.com/v1ack/linkShorter/pkg"
)

func (t *testSuite) TestShorterGet() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	link := "https://googl1r2.com"
	shortLink := hashString(link)
	err = t.store.Create(ctx, link, shortLink)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow(err.Error())
	}

	get, err := service.Get(ctx, &linkShorter.GetLinkRequest{Link: shortLink})
	if err != nil {
		t.FailNow(err.Error())
	}
	t.Equal(get.OriginalLink, link)
}

func (t *testSuite) TestShorterGet_NotCreated() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	_, err = t.store.Get(ctx, "00000000")
	t.ErrorIs(err, store.ErrNotFound)

	_, err = service.Get(ctx, &linkShorter.GetLinkRequest{Link: "00000000"})
	t.EqualError(err, "rpc error: code = NotFound desc = Not found")
}

func (t *testSuite) TestShorterGet_BadRequest() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	_, err = service.Get(ctx, &linkShorter.GetLinkRequest{})
	t.EqualError(err, "rpc error: code = InvalidArgument desc = Bad request")
}
