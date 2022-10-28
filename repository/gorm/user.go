package gorm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) RegisterUser(user *models.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.Message, "username") {
					err = fmt.Errorf("User with username %s already exists.", user.Username)
				} else if strings.Contains(pgErr.Message, "email") {
					err = fmt.Errorf("User with email %s already exists.", user.Email)
				}
				return err
			}
		}
		return err
	}
	return err
}
