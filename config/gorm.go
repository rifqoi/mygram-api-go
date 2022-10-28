package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jusidama18/mygram-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost string
	dbUser string
	dbPass string
	dbName string
	dbPort string
)

func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")

}

func ConnectPostgresGORM() (*gorm.DB, error) {
	getEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(&models.User{})

	return db, nil
}
