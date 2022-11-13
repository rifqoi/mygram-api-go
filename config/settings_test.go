package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetEnv_Success(t *testing.T) {
	os.Setenv("TEST_ENV_FILE", "../.env")
	DB_HOST := GetEnv("DB_HOST")

	assert.NotNil(t, DB_HOST)
	assert.Equal(t, "localhost", DB_HOST)
}

func TestGetEnv_Failed(t *testing.T) {
	os.Setenv("TEST_ENV_FILE", "../.env")
	UNKNOWN_ENV := GetEnv("UNKNOWN_ENV")

	assert.Equal(t, "", UNKNOWN_ENV)
}
