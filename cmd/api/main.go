package main

import (
	"cleanArch/todos/config"
	"cleanArch/todos/db"
	"cleanArch/todos/services/server"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting api server")
	//
	cfg := config.NewConfig()
	db := db.GetPostGresInstace(cfg, true)
	s := server.NewServer(cfg, db, logrus.New(), nil)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
