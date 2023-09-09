package services

import "backend/models"

type ITodoService interface {
	GetAll() ([]models.Todo, error)
	Update(int, models.Todo) (models.Todo, error)
	GetById(int) (models.Todo, error)
	Save(todo models.Todo) (models.Todo, error)
	Delete(int) error
	GetUserTodos(int) ([]models.Todo, error)
}
