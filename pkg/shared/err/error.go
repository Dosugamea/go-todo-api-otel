package shared_err

import "errors"

var ErrInvalidId = errors.New("invalid id specified")
var ErrNotFound = errors.New("specified item was not found")
var ErrInternalServerError = errors.New("internal server error")
