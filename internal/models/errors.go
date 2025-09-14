package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: no mtaching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
