package repositories

import "backend/models"

type ITodo interface {
	GetAll() ([]models.Todo, error)
	GetById(int) (models.Todo, error)
	Update(int, models.Todo) (models.Todo, error)
	Save(models.Todo) (models.Todo, error)
	Delete(int) error
	GetByUser(int) ([]models.Todo, error)
}
