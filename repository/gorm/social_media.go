package gorm

import (
	"errors"
	"fmt"

	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) repository.SocialMediaRepository {
	return &socialMediaRepo{
		db: db,
	}
}

func (sm *socialMediaRepo) CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := sm.db.Create(socialMedia).Error
	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (sm *socialMediaRepo) GetAllSocialMedia() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err :=
		sm.db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username")
		}).Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (sm *socialMediaRepo) UpdatesocialMedia(currentsocialMedia, newsocialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := sm.db.Model(&currentsocialMedia).Updates(&newsocialMedia).Find(&currentsocialMedia).Error
	if err != nil {
		return nil, err
	}
	return currentsocialMedia, nil
}

func (sm *socialMediaRepo) DeletesosialMedia(socialMedia *models.SocialMedia) error {
	err := sm.db.Delete(socialMedia).Error
	return err
}

func (sm *socialMediaRepo) FindSocialMediaByID(socialMediaId int) (*models.SocialMedia, error) {
	var socialMedia *models.SocialMedia

	err := sm.db.First(&socialMedia, socialMediaId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("SocialMedia with id %d not found.", socialMediaId)
		}
		return nil, err
	}

	return socialMedia, nil
}
