package repositories

import "backend/models"

type IUser interface {
	Store(user models.User) error
	Get(user models.User) (models.User, error)
}
