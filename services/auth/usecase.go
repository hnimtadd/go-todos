package auth

import (
	"cleanArch/todos/services/model"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, limit int) (*model.User, error)
	SignIn(ctx context.Context, username, password string) (string, error) // return token
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
