package repository

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/model"
	"strings"

	"context"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) auth.UserRepository {
	return &UserGormRepository{db: db}
}
func (ur *UserGormRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := ur.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (ur *UserGormRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where(&model.User{
		Username: strings.ToLower(username),
	}).Limit(1).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *UserGormRepository) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where(&model.User{
		Id: userId,
	}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}
