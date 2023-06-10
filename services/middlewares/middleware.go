package middlewares

import "cleanArch/todos/services/auth"

type MiddlewareManager struct {
	authUC auth.UseCase
}

func NewMiddlewareManager(authUC auth.UseCase) *MiddlewareManager {
	return &MiddlewareManager{
		authUC: authUC,
	}
}
