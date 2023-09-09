package services

import "backend/models"

type IUserService interface {
	Login(user models.User) (string, error)
	Register(user models.User) error
}
