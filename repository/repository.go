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

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) (*models.Photo, error)
	GetAllPhotos() ([]models.Photo, error)
	UpdatePhoto(currentPhoto, newPhoto *models.Photo) (*models.Photo, error)
	DeletePhoto(photo *models.Photo) error
	FindPhotoByID(id int) (*models.Photo, error)
}
