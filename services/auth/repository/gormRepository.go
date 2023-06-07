package repository

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/model"
	"strings"

	"context"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) auth.UserRepository {
	return &gormUserRepository{db: db}
}
func (ur *gormUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := ur.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (ur *gormUserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where(&model.User{
		Username: strings.ToLower(username),
	}).Limit(1).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *gormUserRepository) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where(&model.User{
		Id: userId,
	}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}
