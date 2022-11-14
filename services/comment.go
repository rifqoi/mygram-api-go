package services

import (
	"fmt"

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

func (c *CommentService) CreateComment(req parameters.CreateComment, userId int) (*models.Comment, error) {
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

func (c *CommentService) UpdateComment(req parameters.UpdateComment, commentID, userID int) (*models.CommentUpdate, error) {
	currentComment, err := c.repo.FindCommentByID(commentID)
	if err != nil {
		return nil, err
	}

	if currentComment.UserID != userID {
		return nil, fmt.Errorf("Comment with id %d is not a comment owned by user with id %d.", commentID, userID)
	}

	newComment := &models.Comment{Message: req.Message}

	updatedPhoto, err := c.repo.UpdateComment(currentComment, newComment)
	if err != nil {
		return nil, err
	}

	resComment := &models.CommentUpdate{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Photo.Title,
		Caption:   updatedPhoto.Photo.Caption,
		PhotoURL:  updatedPhoto.Photo.PhotoURL,
		UserID:    updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	return resComment, nil
}

func (c *CommentService) DeleteComment(commentID int, userID int) error {
	comment, err := c.repo.FindCommentByID(commentID)
	if err != nil {
		return err

	}

	if comment.UserID != userID {
		return fmt.Errorf("Photo with id %d is not a photo owned by user with id %d.", commentID, userID)
	}

	err = c.repo.DeleteComment(comment)
	if err != nil {
		return fmt.Errorf("Error deleting photo: %v", err)
	}

	return err
}
