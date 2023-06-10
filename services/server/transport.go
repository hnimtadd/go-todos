package server

import (
	authRepository "cleanArch/todos/services/auth/repository"
	authHttp "cleanArch/todos/services/auth/transport/http"
	authUsecase "cleanArch/todos/services/auth/usecase"
	"cleanArch/todos/services/middlewares"

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

	v1 := c.Group("/api/v1")
	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todos")

	mw := middlewares.NewMiddlewareManager(authUC)
	authHttp.MapAuthTransport(authGroup, authTransport)
	todoHttp.MapTodosTransport(todoGroup, todoTransport, mw)

	return nil
}
