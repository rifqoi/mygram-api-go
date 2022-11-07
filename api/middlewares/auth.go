package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/helpers"
	"github.com/jusidama18/mygram-api-go/services"
)

type Middleware struct {
	userService *services.UserService
}

func NewMiddleware(userService *services.UserService) *Middleware {
	return &Middleware{
		userService: userService,
	}
}

func (m *Middleware) Authorization(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		responses.UnauthorizedRequest(c, "JWT Token must be provided.")
		return
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	claims, err := helpers.ValidateToken(tokenString)
	// Cek apakah "id" exist didalam map claims
	if _, exist := claims["id"]; !exist {
		responses.UnauthorizedRequest(c, "Id not exists!")
		return
	}

	// Default unmarshal dari encoding/json itu ke float64 sehingga di cast ke float64
	// https://stackoverflow.com/questions/70705673/panic-interface-conversion-interface-is-float64-not-int64
	id := int(claims["id"].(float64))
	user, err := m.userService.FindUserByID(id)
	if err != nil {
		responses.UnauthorizedRequest(c, err.Error())
		return
	}

	c.Set("claims", claims)
	c.Set("userInfo", user)
	c.Next()
}
