package http

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/auth/presenter"
	"cleanArch/todos/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authTransport struct {
	authUc auth.UseCase
}

func NewAuthTransport(authUc auth.UseCase) auth.AuthTransport {
	return &authTransport{
		authUc: authUc,
	}
}

func (at *authTransport) SignIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &presenter.SignInRequest{}
		if err := utils.ReadRequest(ctx, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if input.Username == "" || input.Password == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		token, err := at.authUc.SignIn(ctx.Request().Context(), input.Username, input.Password)
		if err != nil {
			if err == auth.ErrUserNotFount {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			} else if err == auth.ErrWrongPassword {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return ctx.JSON(http.StatusOK, presenter.SignInResponse{Token: token})
	}

}
func (at *authTransport) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &presenter.SignUpRequest{}
		if err := utils.ReadRequest(ctx, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if input.Username == "" || input.Password == "" || input.Limit < 0 {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		user, err := at.authUc.SignUp(ctx.Request().Context(), input.Username, input.Password, input.Limit)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, presenter.SignUpRepsonse{Id: user.Id, Username: user.Username, Limit: user.Limit})
	}
}
