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

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetAllComment() ([]models.Comment, error)
	UpdateComment(currentComment, newComment *models.Comment) (*models.Comment, error)
	DeleteComment(comment *models.Comment) error
	FindCommentByID(id int) (*models.Comment, error)
}

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMedia() ([]models.SocialMedia, error)
	UpdateSocialMedia(currentsocialMedia, newsocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSosialMedia(SocialMedia *models.SocialMedia) error
	FindSocialMediaByID(id int) (*models.SocialMedia, error)
}
