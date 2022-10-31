package repository

import (
	"github.com/jusidama18/mygram-api-go/models"
)

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	Login(email string, password string) (*models.User, error)
	UpdateUser(currentUser *models.User, updatedUser *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	FindUserByID(id int) (*models.User, error)
}
