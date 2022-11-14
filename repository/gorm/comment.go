package gorm

import (
	"errors"
	"fmt"

	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Create(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, err
}

func (c *commentRepository) GetAllComment() ([]models.Comment, error) {
	var comment []models.Comment

	err := c.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Find(&comment).Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepository) UpdateComment(currentComment, newComment *models.Comment) (*models.Comment, error) {
	err := c.db.Model(&currentComment).Updates(&newComment).Find(&currentComment).Error
	if err != nil {
		return nil, err
	}
	return currentComment, nil
}

func (c *commentRepository) DeleteComment(comment *models.Comment) error {
	err := c.db.Delete(comment).Error
	return err
}

func (c *commentRepository) FindCommentByID(id int) (*models.Comment, error) {
	var comment *models.Comment

	err := c.db.First(&comment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Comment with id %d not found.", id)
		}
		return nil, err
	}

	return comment, nil
}
