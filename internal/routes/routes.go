package routes

import "github.com/gofiber/fiber/v2"

type Route struct {
	App fiber.Router
}

func NewRoute(app fiber.Router) *Route {
	return &Route{
		App: app,
	}
}

func (r *Route) Setup() {
	AuthRoutes(r.App)
}
