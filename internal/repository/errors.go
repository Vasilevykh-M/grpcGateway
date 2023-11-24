package repository

import "errors"

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrNetwork        = errors.New("err network")
)
