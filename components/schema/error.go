package schema

import "errors"

var (
	ErrUsernameNoMatch = errors.New("user not found")
	ErrSource          = errors.New("cant access resource")
	ErrUrlFormat       = errors.New("wrong url request format")
	ErrSourceNotFound  = errors.New("source not yet avaible")
)
