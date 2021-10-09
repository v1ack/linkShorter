package app

import (
	"context"
	linkShorter "github.com/v1ack/linkShorter/pkg"
)

func (t *testSuite) TestShorterCreate() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	const link = "https://yandex.ru"
	create, err := service.Create(ctx, &linkShorter.CreateLinkRequest{Link: link})
	if err != nil {
		t.FailNow(err.Error())
	}
	t.Equal(create.ShortLink, "xd9wY9MNLv")

	get, err := t.store.Get(ctx, create.ShortLink)
	if err != nil {
		t.FailNow(err.Error())
	}
	t.Equal(get, link)
}

func (t *testSuite) TestShorterCreate_Twice() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	const link = "https://github.com"
	create, err := service.Create(ctx, &linkShorter.CreateLinkRequest{Link: link})
	if err != nil {
		t.FailNow(err.Error())
	}
	t.Equal(create.ShortLink, "cXMAQGM2CG")

	get, err := t.store.Get(ctx, create.ShortLink)
	if err != nil {
		t.FailNow(err.Error())
	}
	t.Equal(get, link)
}

func (t *testSuite) TestShorterCreate_BadRequest() {
	ctx := context.Background()
	service, err := NewService(t.store)
	if err != nil {
		t.FailNow(err.Error())
	}

	_, err = service.Create(ctx, &linkShorter.CreateLinkRequest{})
	t.EqualError(err, "rpc error: code = InvalidArgument desc = Bad request")
}
