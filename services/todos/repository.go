package todos

import (
	"cleanArch/todos/services/model"
	"context"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	CountTodo(ctx context.Context, userId string) (int, error)
}
