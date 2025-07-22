package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Route) UserRoutes(app fiber.Router) {

	app.Get("/reset-password", func(c *fiber.Ctx) error {
		return c.SendString("Reset Password endpoint")
	})

	app.Get("/verify-email", func(c *fiber.Ctx) error {
		return c.SendString("Verify Email endpoint")
	})

	app.Delete("/delete-account", func(c *fiber.Ctx) error {
		return c.SendString("Delete Account endpoint")
	})

	app.Patch("/change-password", func(c *fiber.Ctx) error {
		return c.SendString("Change Password endpoint")
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return c.SendString("Update User endpoint")
	})
}
