package todos

import (
	"cleanArch/todos/services/model"
	"context"
)

type UseCase interface {
	CreateTodo(ctx context.Context, userId, content string) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
}
