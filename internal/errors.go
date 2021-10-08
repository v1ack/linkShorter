package internal

import "errors"

//type NotFound error
//type Unavailable error
//type AlreadyExists error

var ErrNotFound = errors.New("provider: key not found")
var ErrUnavailable = errors.New("provider: unavailable")
var ErrAlreadyExists = errors.New("provider: already exists")
