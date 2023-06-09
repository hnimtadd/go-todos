package http

import (
	"cleanArch/todos/services/auth"

	"github.com/labstack/echo/v4"
)

func MapAuthTransport(authGroup *echo.Group, tp auth.AuthTransport) {
	authGroup.POST("/register", tp.SignUp())
	authGroup.POST("/login", tp.SignIn())
}
