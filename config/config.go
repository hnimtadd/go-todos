package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port                  string `env:"PORT" envDefault:"8080"`
	HashSalt              string `env:"HASH_SALT,required"`
	SigningKey            string `env:"SIGNING_KEY,required"`
	TokenTTL              int64  `env:"TOKEN_TTL,required"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	DatabaseConnectionURL string `env:"CONNECTION_URL,required"`
}

func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)
	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	return &cfg
}
