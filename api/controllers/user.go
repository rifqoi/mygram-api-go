package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/services"
)

type UserController struct {
	svc *services.UserService
}

func NewUser(svc *services.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var req parameters.UserRegister

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	errs := req.Validate()
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := u.svc.RegisterUser(&req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusCreated, user, "user successfully created")
}
