package services

import (
	"backend/helpers"
	"backend/models"
	"backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.IUser
}

func (us *UserService) Login(reqBody models.User) (string, error) {
	user, err := us.repo.Get(reqBody)
	if err != nil {
		return "", err
	}
	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		return "", &ErrClient{msg: "Invalid password"}
	}
	return helpers.GenerateToken(user.Username, int(user.ID))
}

func (us *UserService) Register(reqBody models.User) error {
	var newUser models.User
	if reqBody.Username == "" || reqBody.Password == "" {
		return &ErrClient{msg: "fill all fields"}
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 5)
	newUser.Username = reqBody.Username
	newUser.Password = string(hashedPassword)
	err := us.repo.Store(newUser)
	return err
}

func NewUserService(UserRepository repositories.IUser) IUserService {
	return &UserService{
		repo: UserRepository,
	}
}

// Errors section
type ErrClient struct {
	msg string
}

func (e *ErrClient) Error() string {
	return e.msg
}
