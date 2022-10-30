package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func (m *Middleware) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			responses.UnauthorizedRequest(c, "JWT Token must be provided")
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := helpers.ValidateToken(tokenString)
		if err != nil {
			responses.UnauthorizedRequest(c, err.Error())
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			responses.UnauthorizedRequest(c, "Invalid token.")
			return
		}

		email := claims["email"].(string)
		user, err := m.userService.FindUserByEmail(email)
		if err != nil {
			responses.UnauthorizedRequest(c, err.Error())
		}

		c.Set("claims", claims)
		c.Set("userInfo", user)
	}
}
