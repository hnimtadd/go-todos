package auth

import (
	"cleanArch/todos/services/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}
