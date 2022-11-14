package services

import (
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c *CommentService) CreateComment(req parameters.Comment, userId int) (*models.Comment, error) {
	comment := &models.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  userId,
	}

	resComment, err := c.repo.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return resComment, nil
}

func (c *CommentService) GetAllComment() ([]models.CommentGetAll, error) {
	comments, err := c.repo.GetAllComment()
	if err != nil {
		return nil, err
	}
	var resComment []models.CommentGetAll
	for _, comment := range comments {
		newComment := models.CommentGetAll{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			User: models.CommentUser{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: models.CommentPhoto{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   comment.Photo.UserID,
			},
		}
		resComment = append(resComment, newComment)
	}
	return resComment, nil
}
