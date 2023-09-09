package main

import (
	"backend/controllers"
	"backend/infrastructure"
	"backend/middlewares"
	"backend/repositories"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	DB := infrastructure.ConnectToDB()

	todoRepo := repositories.NewTodoRepository(DB)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	userRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := app.Group("api")
	todoRoutes := api.Group("todos")
	todoRoutes.Get("", todoController.IndexTodos)
	todoRoutes.Get("/user/:id", middlewares.AuthMiddleware, todoController.GetUserTodos)
	todoRoutes.Get("/:id", todoController.GetTodo)
	todoRoutes.Post("", todoController.CreateTodo)
	todoRoutes.Put("/:id", todoController.UpdateTodo)
	todoRoutes.Delete("/:id", todoController.DeleteTodo)

	users := api.Group("/users")
	users.Post("/register", userController.Register)
	users.Post("/login", userController.Login)

	app.Listen(":8082")
}
