package utils

import "errors"

var ErrInternalServerError = errors.New("internal server error")
var ErrIncorrectPassword = errors.New("incorrect password")
var ErrUserNotFound = errors.New("user not found")
