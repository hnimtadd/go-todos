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
	ctx    echo.Context
}

func NewAuthTransport(authUc auth.UseCase, ctx echo.Context) auth.AuthTransport {
	return &authTransport{
		authUc: authUc,
		ctx:    ctx,
	}
}

func (at *authTransport) SignIn() error {
	input := &presenter.SignInRequest{}
	if err := utils.ReadRequest(at.ctx, input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	token, err := at.authUc.SignIn(at.ctx.Request().Context(), input.Username, input.Password)
	if err != nil {
		if err == auth.ErrUserNotFount {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if err == auth.ErrWrongPassword {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return at.ctx.JSON(http.StatusOK, presenter.SignInResponse{Token: token})

}
func (at *authTransport) SignUp() error {
	input := &presenter.SignUpRequest{}
	if err := utils.ReadRequest(at.ctx, input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	user, err := at.authUc.SignUp(at.ctx.Request().Context(), input.Username, input.Password, input.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return at.ctx.JSON(http.StatusCreated, presenter.SignUpRepsonse{Id: user.Id, Username: user.Username, Limit: user.Limit})
}
