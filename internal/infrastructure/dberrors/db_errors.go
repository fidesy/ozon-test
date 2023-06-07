package dberrors

import "errors"

var (
	ErrInvalidDatabaseName = errors.New("invalid database name, available options: postgres, in-memory")
	ErrHashDoesNotExist    = errors.New("hash does not exist")
)
