package helpers

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	token string
}

func TestJWTTestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) SetupTest() {
	os.Setenv("TEST_ENV_FILE", "../.env")
}

func (s *Suite) TestGenerateToken_Success() {
	email := "hacktiv8@gmail.com"
	id := 1
	token, err := GenerateToken(email, id)

	s.NotNil(token)
	s.Nil(err)
}

func (s *Suite) TestValidateToken_Success() {
	email := "hacktiv8@gmail.com"
	id := 1
	token, err := GenerateToken(email, id)
	s.Require().NoError(err)
	s.Require().NotNil(token)

	claims, err := ValidateToken(token)
	s.Nil(err)
	s.NotNil(claims)

	s.NotNil(claims["id"])
	s.Equal(float64(id), claims["id"])

	s.NotNil(claims["email"])
	s.Equal(email, claims["email"])

	s.NotNil(claims["exp"])

	exp := int64(claims["exp"].(float64))
	expiredAt := time.Unix(exp, 0)

	isExp := time.Now().After(expiredAt)
	s.False(isExp)
}

func (s *Suite) TestValidateToken_InvalidMethod() {
	token := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg5NTExODgsImlhdCI6MTY2ODM0NjM4OCwiZW1haWwiOiJoOEBoOC5jb216eCIsImlkIjo2NX0.Pj7Kvw0hINJeZajKbNU-qGz7Hy7c6wcOjotevkVJmHHGeooQxBDxCFN9q2DNyATH"
	expError := fmt.Errorf("Invalid signing method.")

	claims, err := ValidateToken(token)

	s.Nil(claims)
	s.NotNil(err)
	s.Equal(expError.Error(), err.Error())
}
