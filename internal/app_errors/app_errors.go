package apperrors

import "errors"

var (
	ErrURLNotFound          = errors.New("url not found")
	ErrURLAlreadyExists     = errors.New("url already exists")
	ErrAliasAlreadyOccupied = errors.New("alias already occupied")
)
