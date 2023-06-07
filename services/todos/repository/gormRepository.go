package repository

import (
	"cleanArch/todos/services/model"
	"cleanArch/todos/services/todos"
	"context"

	"gorm.io/gorm"
)

type todoGormRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todos.TodoRepository {
	return &todoGormRepository{db: db}
}

func (tr *todoGormRepository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	res := tr.db.WithContext(ctx).Create(&todo)
	if res.Error != nil {
		return res.Error
	}
	return nil

}
func (tr *todoGormRepository) GetTodosByUserId(ctx context.Context, userId string) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := tr.db.WithContext(ctx).Where(&model.Todo{CreatedBy: userId}).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (tr *todoGormRepository) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := tr.db.WithContext(ctx).Find(&todos).Limit(200).Error
	if err != nil {
		return nil, err
	}
	return todos, nil

}
func (tr *todoGormRepository) CountTodo(ctx context.Context, userId string) (int, error) {
	var res int64
	err := tr.db.WithContext(ctx).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return int(res), nil

}
