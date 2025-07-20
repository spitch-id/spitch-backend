package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/internal/routes"
)

type ServerConfig struct {
	App fiber.Router
}

func NewServerConfig(sc *ServerConfig) {
	versionOne := sc.App.Group("/v1")

	routeConfig := routes.NewRoute(versionOne)
	routeConfig.Setup()
}
