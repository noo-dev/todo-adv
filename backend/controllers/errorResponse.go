package controllers

import "github.com/gofiber/fiber/v2"

func errorResponse(c *fiber.Ctx, status int, msg string) error {

	return c.Status(status).JSON(fiber.Map{
		"error": msg,
	})
}
