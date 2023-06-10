package usecase

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/model"
	"cleanArch/todos/services/todos"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type todoUsecase struct {
	todoRepo todos.TodoRepository
	userRepo auth.UserRepository
}

func NewTodoUsecase(todoRepo todos.TodoRepository, userRepo auth.UserRepository) todos.UseCase {
	return &todoUsecase{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}

func (tu todoUsecase) CreateTodo(ctx context.Context, userId, content string) error {
	todo := &model.Todo{
		Id:        uuid.New().String(),
		Content:   content,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}
	count, err := tu.todoRepo.CountTodo(ctx, userId)

	if err != nil {
		return err
	}
	fmt.Println("hrllo")

	user, err := tu.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	// check if user todo is react limit of this account
	if user.Limit > count {
		return tu.todoRepo.CreateTodo(ctx, todo)
	} else {
		return errors.New("Limit exeeded")
	}

}
func (tu todoUsecase) GetTodosByUserId(ctx context.Context, userId string) ([]*model.Todo, error) {
	todos, err := tu.todoRepo.GetTodosByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return todos, nil

}
func (tu todoUsecase) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := tu.todoRepo.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil

}
