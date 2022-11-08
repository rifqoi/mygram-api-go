package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/models"
)

func GetUser(c *gin.Context) (*models.User, error) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		return nil, fmt.Errorf("context error")
	}

	user := userInfo.(*models.User)
	return user, nil
}
