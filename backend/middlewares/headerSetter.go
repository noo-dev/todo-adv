package middlewares

import "github.com/gofiber/fiber/v2"

func SetCustomHeader(c *fiber.Ctx) error {
	c.Set("Custom-Header", "custom value")
	return c.Next()
}
