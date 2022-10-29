package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/helpers"
	"github.com/jusidama18/mygram-api-go/services"
)

func Authorization(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			responses.BadRequestError(c, "Invalid token")
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := helpers.ValidateToken(tokenString)
		if err != nil {
			responses.BadRequestError(c, "Invalid token")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			responses.UnauthorizedRequest(c)
			return
		}

		email := claims["email"].(string)
		user, err := userService.FindUserByEmail(email)
		if err != nil {
			responses.UnauthorizedRequest(c)
		}

		c.Set("claims", claims)
		c.Set("userInfo", user)
	}
}
