package services

import (
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
