package store

import "errors"

var ErrNotFound = errors.New("provider: key not found")
var ErrUnavailable = errors.New("provider: unavailable")
var ErrAlreadyExists = errors.New("provider: already exists")
