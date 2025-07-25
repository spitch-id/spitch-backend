package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/spitch-id/spitch-backend/internal/entity"
)

func NewEnv() *entity.Env {
	env := entity.Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the .env file: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.SERVER_APP_ENV == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
