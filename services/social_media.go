package services

import (
	"fmt"

	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
)

type SocialMediaService struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		repo: repo,
	}
}

func (sm *SocialMediaService) CreateSocialMedia(req parameters.SocialMediaCreate, userID int) (*models.SocialMedia, error) {
	socialMedia := &models.SocialMedia{
		UserID:         userID,
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaURL,
	}
	resp, err := sm.repo.CreateSocialMedia(socialMedia)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (sm *SocialMediaService) GetAllSocialMedia() ([]models.SocialMediaGet, error) {

	socialMedias, err := sm.repo.GetAllSocialMedia()
	if err != nil {
		return nil, err
	}

	parsedSocialMedia := parseSocialMediaGet(socialMedias)
	return parsedSocialMedia, nil
}

func (sm *SocialMediaService) UpdateSocialMedia(req parameters.SocialMediaUpdate, smID int, userID int) (*models.SocialMediaUpdate, error) {
	currentSocialMedia, err := sm.repo.FindSocialMediaByID(smID)
	if err != nil {
		return nil, err
	}

	if currentSocialMedia.UserID != userID {
		return nil, fmt.Errorf("Social Media with id %d is not owned by user with id %d", smID, userID)
	}
	newSocialMedia := &models.SocialMedia{
		SocialMediaUrl: req.SocialMediaURL,
		Name:           req.Name,
	}

	updatedSocialMedia, err := sm.repo.UpdateSocialMedia(currentSocialMedia, newSocialMedia)
	if err != nil {
		return nil, err
	}

	resp := &models.SocialMediaUpdate{
		ID:             updatedSocialMedia.ID,
		Name:           updatedSocialMedia.Name,
		SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
		UserID:         updatedSocialMedia.UserID,
		UpdatedAt:      currentSocialMedia.UpdatedAt,
	}

	return resp, nil
}

func (sm *SocialMediaService) DeleteSocialMedia(smID, userID int) error {
	currentSocialMedia, err := sm.repo.FindSocialMediaByID(smID)
	if err != nil {
		return err
	}

	if currentSocialMedia.UserID != userID {
		return err
	}

	err = sm.repo.DeleteSosialMedia(currentSocialMedia)
	if err != nil {
		return fmt.Errorf("Error deleting social media: %v", err)
	}

	return nil
}

func parseSocialMediaGet(socialMedias []models.SocialMedia) []models.SocialMediaGet {
	var parsedSocialMedia []models.SocialMediaGet
	for _, sm := range socialMedias {
		newSM := models.SocialMediaGet{
			ID:             sm.ID,
			Name:           sm.Name,
			SocialMediaUrl: sm.SocialMediaUrl,
			UserID:         sm.UserID,
			User: models.SocialMediaUserGet{
				Email:    sm.User.Email,
				Username: sm.User.Username,
			},
			CreatedAt: sm.CreatedAt,
			UpdatedAt: sm.UpdatedAt,
		}
		parsedSocialMedia = append(parsedSocialMedia, newSM)
	}
	return parsedSocialMedia
}
