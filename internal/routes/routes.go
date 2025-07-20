package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spitch-id/spitch-backend/domain"
)

type Route struct {
	App         fiber.Router
	Validate    *validator.Validate
	UserHandler domain.UserHandler
}

func NewRoute(app fiber.Router, validator validator.Validate, userHandler domain.UserHandler) *Route {
	return &Route{
		App:         app,
		Validate:    &validator,
		UserHandler: userHandler,
	}
}

func (r *Route) Setup() {
	versionOne := r.App.Group("/v1")
	r.AuthRoutes(versionOne)
}
