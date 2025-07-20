package routes

import "github.com/gofiber/fiber/v2"

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login endpoint")
	})

	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("Register endpoint")
	})

	auth.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("Logout endpoint")
	})

	auth.Get("/forgot-password", func(c *fiber.Ctx) error {
		return c.SendString("Forgot Password endpoint")
	})

	auth.Get("/verify-email", func(c *fiber.Ctx) error {
		return c.SendString("Verify Email endpoint")
	})
}
