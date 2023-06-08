package todos

import "github.com/labstack/echo/v4"

type TodoTransport interface {
	GetAll() echo.HandlerFunc
	GetUserTodos() echo.HandlerFunc
	AddTodo() echo.HandlerFunc
}
