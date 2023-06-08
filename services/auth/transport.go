package auth

import "github.com/labstack/echo/v4"

type AuthTransport interface {
	SignIn() echo.HandlerFunc
	SignUp() echo.HandlerFunc
}
