package http

import (
	"cleanArch/todos/services/todos"

	"github.com/labstack/echo/v4"
)

func MapTodosTransport(todoGroup *echo.Group, tp todos.TodoTransport) {
	todoGroup.POST("/create", tp.AddTodo())
	todoGroup.GET("/", tp.GetAll())
	todoGroup.POST("/:userId", tp.GetUserTodos())
}
