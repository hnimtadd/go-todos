package auth

import "errors"

var (
	ErrUserNotFount       = errors.New("user not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErruserExisted        = errors.New("user existsed")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
