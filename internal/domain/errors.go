package domain

import "errors"

var (
	ErrURLIsInvalid     = errors.New("the URL provided is invalid and cannot be processed")
	ErrURLNotFound      = errors.New("URL not found")
	ErrURLAlreadyExists = errors.New("URL already exists")
)
