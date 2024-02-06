package shared_err

import "errors"

var ErrInvalidId = errors.New("invalid id specified")
var ErrInternalServerError = errors.New("internal server error")
