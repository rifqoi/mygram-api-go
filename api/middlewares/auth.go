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

		if !strings.HasPrefix(authHeader, "Bearer ") {
			responses.UnauthorizedRequest(c, "JWT Token must be provided.")
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

		// Cek apakah "id" exist didalam map claims
		if _, exist := claims["id"]; !exist {
			responses.UnauthorizedRequest(c, "Invalid token")
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
	}
}
