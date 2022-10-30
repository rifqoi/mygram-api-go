package repository

import "github.com/jusidama18/mygram-api-go/models"

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	Login(email string, password string) error
}
