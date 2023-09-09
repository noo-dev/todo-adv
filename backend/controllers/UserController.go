package controllers

import (
	"backend/models"
	"backend/services"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.IUserService
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	err = uc.userService.Register(user)
	if err != nil {
		var status int
		if errors.Is(err, &services.ErrClient{}) {
			status = 400
		} else {
			status = 500
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "user was registered successfully",
	})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	token, err := uc.userService.Login(user)
	if err != nil {
		var status int
		if errors.Is(err, &services.ErrClient{}) {
			status = 400
		} else {
			status = 500
		}
		return c.Status(status).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"token":   token,
		"success": true,
		"message": "you are successfully logged in",
	})
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{
		userService: service,
	}
}
