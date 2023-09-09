package repositories

import (
	"backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Store(userBody models.User) error {
	tx := repo.DB.Save(&userBody)
	return tx.Error
}

func (repo *UserRepository) Get(credentials models.User) (models.User, error) {
	var user models.User
	err := repo.DB.Find(&user, "username = ?", credentials.Username).Error
	if err != nil || user.ID == 0 {
		return models.User{}, err
	}
	return user, nil
}

func NewUserRepository(DB *gorm.DB) IUser {
	return &UserRepository{DB: DB}
}
