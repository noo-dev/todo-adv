package repositories

import (
	"backend/models"
	"fmt"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (repo *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	tx := repo.DB.Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todos, nil
}

func (repo *TodoRepository) GetById(id int) (models.Todo, error) {
	var todo models.Todo
	tx := repo.DB.Find(&todo, "id = ?", id)
	if tx.Error != nil {
		return todo, tx.Error
	}
	if todo.ID == 0 {
		return todo, &NotFoundError{id}
	}
	return todo, nil
}

func (repo *TodoRepository) GetByUser(userID int) ([]models.Todo, error) {
	var todos []models.Todo
	err := repo.DB.Find(&todos, "user_id = ?", userID).Error
	return todos, err
}

func (repo *TodoRepository) Save(todo models.Todo) (models.Todo, error) {
	err := repo.DB.Create(&todo).Error
	return todo, err
}

func (repo *TodoRepository) Delete(id int) error {
	todo := models.Todo{}
	tx := repo.DB.Delete(&todo, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *TodoRepository) Update(id int, todoBody models.Todo) (models.Todo, error) {
	var todo models.Todo
	tx := repo.DB.First(&todo, "id = ?", id)
	if tx.Error != nil {
		return todo, tx.Error
	}
	if todo.ID == 0 {
		return todo, &NotFoundError{id}
	}

	err := repo.DB.Model(&todo).Updates(todoBody).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func NewTodoRepository(DB *gorm.DB) *TodoRepository {
	return &TodoRepository{
		DB: DB,
	}
}

type NotFoundError struct {
	id int
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("Item with id = %d not found", n.id)
}
