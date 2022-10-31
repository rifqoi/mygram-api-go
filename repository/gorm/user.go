package gorm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/repository"
	"golang.org/x/crypto/bcrypt"
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
	err = checkDuplicate(err, user)
	return err
}

func (u *userRepo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := u.db.Where("email", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("User not found")
	}

	return &user, result.Error
}

func (u *userRepo) Login(email string, password string) (*models.User, error) {

	user, err := u.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *userRepo) UpdateUser(currentUser *models.User, updatedUser *models.User) (*models.User, error) {
	err := u.db.Model(&currentUser).Updates(&updatedUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User not found")
		}
		err = checkDuplicate(err, updatedUser)
		return nil, err

	}
	responseUser, err := u.FindUserByEmail(updatedUser.Email)
	if err != nil {
		fmt.Println("gagal find user")
		return nil, err
	}
	return responseUser, nil
}

func (u *userRepo) DeleteUser(user *models.User) error {
	err := u.db.Delete(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("User not found")
	}
	return nil
}

func (u *userRepo) FindUserByID(id int) (*models.User, error) {
	var user models.User
	result := u.db.Where("id", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("User not found")
	}

	return &user, result.Error
}

func checkDuplicate(err error, user *models.User) error {
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
	return nil
}
