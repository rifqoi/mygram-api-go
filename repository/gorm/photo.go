package gorm

import (
	"errors"
	"fmt"

	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repository.PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (p *photoRepository) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	err := p.db.Create(photo).Error
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (p *photoRepository) GetAllPhotos() ([]models.Photo, error) {
	var photos []models.Photo

	err :=
		p.db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username")
		}).Find(&photos).Error

	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *photoRepository) UpdatePhoto(currentPhoto, newPhoto *models.Photo) (*models.Photo, error) {
	err := p.db.Model(&currentPhoto).Updates(&newPhoto).Find(&currentPhoto).Error
	if err != nil {
		return nil, err
	}
	return currentPhoto, nil
}

func (p *photoRepository) DeletePhoto(photo *models.Photo) error {
	err := p.db.Delete(photo).Error
	return err
}

func (p *photoRepository) FindPhotoByID(photoId int) (*models.Photo, error) {
	var photo *models.Photo

	err := p.db.First(&photo, photoId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Photo with id %d not found.", photoId)
		}
		return nil, err
	}

	return photo, nil
}
