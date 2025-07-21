package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	SERVER_APP_NAME                  string `mapstructure:"SERVER_APP_NAME"`
	SERVER_SERVER_NAME               string `mapstructure:"SERVER_SERVER_NAME"`
	SERVER_PORT                      string `mapstructure:"SERVER_PORT"`
	SERVER_HOST                      string `mapstructure:"SERVER_HOST"`
	SERVER_PREFORK                   bool   `mapstructure:"SERVER_PREFORK"`
	SERVER_APP_ENV                   string `mapstructure:"SERVER_APP_ENV"`
	DATABASE_HOST                    string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT                    string `mapstructure:"DATABASE_PORT"`
	DATABASE_USER                    string `mapstructure:"DATABASE_USER"`
	DATABASE_PASS                    string `mapstructure:"DATABSE_PASS"`
	DATABASE_NAME                    string `mapstructure:"DATABASE_NAME"`
	DATABASE_TIMEZONE                string `mapstructure:"DATABASE_TIMEZONE"`
	DATABASE_SSLMODE                 string `mapstructure:"DATABASE_SSLMODE"`
	DATABASE_POOL_IDLE               int    `mapstructure:"DATABASE_POOL_IDLE"`
	DATABASE_MAX_CONNECTIONS         int    `mapstructure:"DATABASE_MAX_CONNECTIONS"`
	DATABASE_MAXLIFETIME_CONNECTIONS int    `mapstructure:"DATABASE_MAXLIFETIME_CONNECTIONS"`
	LOG_LEVEL                        int    `mapstructure:"LOG_LEVEL"`
	ALLOWED_ORIGINS                  string `mapstructure:"ALLOWED_ORIGINS"`
}

func NewEnv() *Env {
	env := Env{}
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
