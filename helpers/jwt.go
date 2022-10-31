package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jusidama18/mygram-api-go/config"
	"github.com/jusidama18/mygram-api-go/models"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNING_KEY = []byte(config.GetEnv("JWT_SIGNING_KEY"))

func GenerateToken(user *models.User) (string, error) {
	LOGIN_EXPIRATION_DURATION := time.Now().Add(time.Hour * 24 * 7).Unix()

	claim := jwt.MapClaims{}
	claim["email"] = user.Email
	claim["id"] = user.ID
	claim["exp"] = LOGIN_EXPIRATION_DURATION
	claim["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claim)

	// SignedString harus type []byte
	signedToken, err := token.SignedString(JWT_SIGNING_KEY)

	return signedToken, err
}

func ValidateToken(encToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method.")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Invalid signing method.")
		}
		return JWT_SIGNING_KEY, nil
	})
	return token, err
}
