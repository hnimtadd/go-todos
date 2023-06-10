package http

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/model"
	"cleanArch/todos/services/todos"
	"cleanArch/todos/services/todos/presenter"
	"cleanArch/todos/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type todoTransport struct {
	todoUc todos.UseCase
}

func NewTodoTranposrt(todoUC todos.UseCase) todos.TodoTransport {
	return &todoTransport{
		todoUc: todoUC,
	}
}
func (tt *todoTransport) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		todos, err := tt.todoUc.GetAllTodos(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, mapTodos(todos))
	}

}
func (tt *todoTransport) GetUserTodos() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Get(auth.CtxUserKey)
		fmt.Println(userId)
		todos, err := tt.todoUc.GetTodosByUserId(ctx.Request().Context(), fmt.Sprintf("%v", userId))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, mapTodos(todos))
	}
}
func (tt *todoTransport) AddTodo() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Get(auth.CtxUserKey)
		request := &presenter.TodoRequest{}
		if err := utils.ReadRequest(ctx, request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := tt.todoUc.CreateTodo(ctx.Request().Context(), fmt.Sprintf("%v", userId), request.Content); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, nil)
	}

}

func mapTodos(todos []*model.Todo) []*presenter.TodoResponse {
	out := make([]*presenter.TodoResponse, len(todos))
	for index, todo := range todos {
		out[index] = mapTodo(todo)
	}
	return out
}

func mapTodo(todo *model.Todo) *presenter.TodoResponse {
	return &presenter.TodoResponse{
		Id:        todo.Id,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt,
		CreatedBy: todo.CreatedBy,
	}
}
