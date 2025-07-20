package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/domain"
	"github.com/spitch-id/spitch-backend/internal/routes"
)

type ServerConfig struct {
	App         fiber.Router
	Validator   *validator.Validate
	UserHandler domain.UserHandler
}

func NewServerConfig(sc *ServerConfig) {
	routeConfig := routes.NewRoute(
		sc.App,
		*sc.Validator,
		sc.UserHandler,
	)
	routeConfig.Setup()
}
