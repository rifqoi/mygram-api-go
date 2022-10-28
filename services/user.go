package services

import (
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUser(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUser(req *parameters.UserRegister) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
		Age:      req.Age,
	}
	err = u.repo.RegisterUser(userModel)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}
