package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if testEnvFile := os.Getenv("TEST_ENV_FILE"); testEnvFile != "" {
		godotenv.Load(testEnvFile)
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	LoadEnv()
	return os.Getenv(key)
}
