package http

import (
	"cleanArch/todos/services/middlewares"
	"cleanArch/todos/services/todos"

	"github.com/labstack/echo/v4"
)

func MapTodosTransport(todoGroup *echo.Group, tp todos.TodoTransport, mw *middlewares.MiddlewareManager) {
	todoGroup.POST("/create", tp.AddTodo(), mw.JWTMiddleware)
	todoGroup.GET("/", tp.GetAll())
	todoGroup.GET("/user", tp.GetUserTodos(), mw.JWTMiddleware)
}
