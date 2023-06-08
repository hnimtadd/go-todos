package server

import (
	authRepository "cleanArch/todos/services/auth/repository"
	authHttp "cleanArch/todos/services/auth/transport/http"
	authUsecase "cleanArch/todos/services/auth/usecase"
	todoRepository "cleanArch/todos/services/todos/repository"
	todoHttp "cleanArch/todos/services/todos/transport/http"
	todoUsecase "cleanArch/todos/services/todos/usecase"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapTransport(c *echo.Echo) error {
	// map repository
	userRepo := authRepository.NewUserRepository(s.db)
	todoRepo := todoRepository.NewTodoRepository(s.db)

	//map usecase
	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	todoUC := todoUsecase.NewTodoUsecase(todoRepo, userRepo)

	// map transport
	authTransport := authHttp.NewAuthTransport(authUC)
	todoTransport := todoHttp.NewTodoTranposrt(todoUC)

	return nil
}
