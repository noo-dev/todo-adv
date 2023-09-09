package services

import (
	"backend/models"
	"backend/repositories"
)

type TodoService struct {
	repo repositories.ITodo
}

func (ts *TodoService) GetAll() ([]models.Todo, error) {
	todos, err := ts.repo.GetAll()
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func (ts *TodoService) Update(id int, todo models.Todo) (models.Todo, error) {
	todo, err := ts.repo.Update(id, todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (ts *TodoService) GetById(id int) (models.Todo, error) {
	todo, err := ts.repo.GetById(id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (ts *TodoService) GetUserTodos(id int) ([]models.Todo, error) {
	todos, err := ts.repo.GetByUser(id)
	return todos, err
}

func (ts *TodoService) Save(todo models.Todo) (models.Todo, error) {
	todo, err := ts.repo.Save(todo)
	return todo, err
}

func (ts *TodoService) Delete(id int) error {
	return ts.repo.Delete(id)
}

func NewTodoService(repo repositories.ITodo) ITodoService {
	return &TodoService{
		repo: repo,
	}
}
