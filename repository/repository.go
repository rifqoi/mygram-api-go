package repository

import "github.com/jusidama18/mygram-api-go/models"

type UserRepository interface {
	RegisterUser(user *models.User) error
}
