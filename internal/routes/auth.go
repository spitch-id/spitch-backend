package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Route) AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login endpoint")
	})

	auth.Post("/register", r.UserHandler.RegisterUser)

	auth.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("Logout endpoint")
	})

	auth.Get("/reset-password", func(c *fiber.Ctx) error {
		return c.SendString("Reset Password endpoint")
	})

	auth.Get("/verify-email", func(c *fiber.Ctx) error {
		return c.SendString("Verify Email endpoint")
	})

	auth.Delete("/delete-account", func(c *fiber.Ctx) error {
		return c.SendString("Delete Account endpoint")
	})

	app.Patch("/change-password", func(c *fiber.Ctx) error {
		return c.SendString("Change Password endpoint")
	})
}
