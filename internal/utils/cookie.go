package utils

import "github.com/gofiber/fiber/v2"

func SetCookie(c *fiber.Ctx, name, value string, maxAge int) error {
	cookie := fiber.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Secure:   true,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}
	c.Cookie(&cookie)
	return nil
}

func GetCookie(c *fiber.Ctx, name string) (string, error) {
	cookie := c.Cookies(name)
	if cookie == "" {
		return "", fiber.NewError(fiber.StatusNotFound, "Cookie not found")
	}
	return cookie, nil
}

func ClearCookie(c *fiber.Ctx, name string) {
	cookie := fiber.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}
	c.Cookie(&cookie)
}
