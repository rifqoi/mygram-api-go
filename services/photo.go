package services

import (
	"fmt"

	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
)

type PhotoService struct {
	repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) *PhotoService {
	return &PhotoService{
		repo: repo,
	}
}

func (p *PhotoService) CreatePhoto(req parameters.CreatePhoto, userId int) (*models.Photo, error) {
	photo := &models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   userId,
	}

	resPhoto, err := p.repo.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	return resPhoto, nil
}

func (p *PhotoService) GetAllPhotos() ([]models.PhotoGet, error) {
	photos, err := p.repo.GetAllPhotos()
	if err != nil {
		return nil, err
	}
	respPhotos := parseGetAllPhotos(photos)

	return respPhotos, nil
}

func (p *PhotoService) UpdatePhoto(req parameters.UpdatePhoto, photoId, userID int) (*models.PhotoUpdate, error) {
	currentPhoto, err := p.repo.FindPhotoByID(photoId)
	if err != nil {
		return nil, err
	}

	if currentPhoto.UserID != userID {
		return nil, fmt.Errorf("Photo with id %d is not a photo owned by user with id %d.", photoId, userID)
	}

	newPhoto := &models.Photo{
		PhotoURL: req.PhotoURL,
		Caption:  req.Caption,
		Title:    req.Caption,
	}

	updatedPhoto, err := p.repo.UpdatePhoto(currentPhoto, newPhoto)
	if err != nil {
		return nil, err
	}

	responsePhoto := parseUpdatePhoto(updatedPhoto)

	return responsePhoto, nil
}

func (p *PhotoService) DeletePhoto(photoID int, userID int) error {
	photo, err := p.repo.FindPhotoByID(photoID)
	if err != nil {
		return err

	}

	if photo.UserID != userID {
		return fmt.Errorf("Photo with id %d is not a photo owned by user with id %d.", photoID, userID)
	}

	err = p.repo.DeletePhoto(photo)
	if err != nil {
		return fmt.Errorf("Error deleting photo: %v", err)
	}

	return err
}

func parseGetAllPhotos(photos []models.Photo) []models.PhotoGet {
	var parsedPhotos []models.PhotoGet
	for _, photo := range photos {
		newPhoto := models.PhotoGet{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.UserID,
			User: models.PhotoUserGet{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		parsedPhotos = append(parsedPhotos, newPhoto)
	}
	return parsedPhotos
}

func parseUpdatePhoto(photo *models.Photo) *models.PhotoUpdate {
	updatedPhoto := &models.PhotoUpdate{
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return updatedPhoto
}
