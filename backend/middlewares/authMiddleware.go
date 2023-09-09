package middlewares

import (
	"backend/helpers"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AuthMiddleware(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, err := strconv.Atoi(userId)
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Error parsing id",
		})
	}
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is empty",
		})
	}
	err = helpers.ValidateToken(token, id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Next()
}
