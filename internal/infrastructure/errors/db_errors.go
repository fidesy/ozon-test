package errors

import "errors"

var (
	ErrInvalidDatabaseName = errors.New("invalid database name, available options: postgres, in-memory")
	ErrHashAlreadyExists   = errors.New("this hash already exists")
	ErrHashDoesNotExist    = errors.New("hash does not exist")
)
