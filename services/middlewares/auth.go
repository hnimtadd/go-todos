package middlewares

import (
	"cleanArch/todos/services/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		userId, err := mw.authUC.ParseToken(ctx.Request().Context(), headerParts[1])
		if err != nil {
			if err == auth.ErrInvalidAccessToken {
				return echo.NewHTTPError(http.StatusUnauthorized)
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		ctx.Set(auth.CtxUserKey, userId)
		return next(ctx)
	}
}
