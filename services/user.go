package services

import (
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/helpers"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUser(req *parameters.UserRegister) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
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

func (u *UserService) Login(email string, password string) (*string, error) {
	user, err := u.repo.Login(email, password)
	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (u *UserService) FindUserByEmail(email string) (*models.User, error) {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) UpdateUser(currentUser *models.User, req *parameters.UserUpdate) (*models.User, error) {
	updatedUser := &models.User{
		Email:    req.Email,
		Username: req.Username,
	}
	responseUser, err := u.repo.UpdateUser(currentUser, updatedUser)
	if err != nil {
		return nil, err
	}
	return responseUser, nil
}

func (u *UserService) DeleteUser(user *models.User) error {
	err := u.repo.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) FindUserByID(id int) (*models.User, error) {
	user, err := u.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
