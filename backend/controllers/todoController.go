package controllers

import (
	"backend/models"
	"backend/repositories"
	"backend/services"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type TodoController struct {
	service services.ITodoService
}

func (tc *TodoController) IndexTodos(c *fiber.Ctx) error {
	todos, err := tc.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"todos":   todos,
	})
}

func (tc *TodoController) UpdateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "id is invalid",
		})
	}
	var todo models.Todo
	err = c.BodyParser(&todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"todos":   err.Error(),
		})
	}
	updatedTodo, err := tc.service.Update(id, todo)
	if err != nil {
		var status int
		if errors.Is(err, &repositories.NotFoundError{}) {
			status = fiber.StatusNotFound
		} else {
			status = fiber.StatusInternalServerError
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"todos":   updatedTodo,
	})
}

func (tc *TodoController) GetTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "id is invalid",
		})
	}

	todo, err := tc.service.GetById(id)
	if err != nil {
		var status int
		if errors.Is(err, &repositories.NotFoundError{}) {
			status = fiber.StatusNotFound
		} else {
			status = fiber.StatusInternalServerError
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"todo":    todo,
	})
}

func (tc *TodoController) CreateTodo(c *fiber.Ctx) error {
	var reqBody models.Todo
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"msg":     "Check your input",
		})
	}
	todo, err := tc.service.Save(reqBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"todo":    todo,
		"msg":     "Todo created successfully",
	})
}

func (tc *TodoController) DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "id is invalid",
		})
	}

	err = tc.service.Delete(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Todo with id = %d deleted successfully", id),
	})
}

func (tc *TodoController) GetUserTodos(c *fiber.Ctx) error {
	userIdParam := c.Params("id")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	todos, err := tc.service.GetUserTodos(userId)
	if err != nil {
		return errorResponse(c, 500, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"todos":   todos,
	})
}

func NewTodoController(service services.ITodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}
