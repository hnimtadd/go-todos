package db

import (
	"cleanArch/todos/config"
	"cleanArch/todos/services/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostGresInstace(cfg *config.Configuration, migrate bool) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if migrate {
		db.AutoMigrate(&model.User{}, &model.Todo{})
	}
	return db

}
