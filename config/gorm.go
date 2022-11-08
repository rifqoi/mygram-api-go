package config

import (
	"fmt"

	"github.com/jusidama18/mygram-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	dbHost := GetEnv("DB_HOST")
	dbUser := GetEnv("DB_USER")
	dbPass := GetEnv("DB_PASS")
	dbName := GetEnv("DB_NAME")
	dbPort := GetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{})

	return db, nil
}
