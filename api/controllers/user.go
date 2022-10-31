package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/models"
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
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(req)
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

func (u *UserController) Login(c *gin.Context) {
	var req parameters.UserLogin

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	token, err := u.svc.Login(req.Email, req.Password)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusOK, gin.H{
		"token": token,
	}, "login successful")
}

func (u *UserController) Check(c *gin.Context) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		responses.InternalServerError(c, fmt.Errorf("context error"))
	}

	user := userInfo.(*models.User)

	responses.SuccessWithData(c, http.StatusOK, user, "success")
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var req parameters.UserUpdate

	// Binding request body ke variabel req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	// Validasi request body apakah sudah benar
	// Variabel errs berbeda dengan err
	errs := parameters.Validate(req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := getUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responseUser, err := u.svc.UpdateUser(user, &req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusOK, responseUser, "user updated successfully")
}

func (u *UserController) DeleteUser(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = u.svc.DeleteUser(user)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Your account has been successfully deleted")
}

func getUser(c *gin.Context) (*models.User, error) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		return nil, fmt.Errorf("context error")
	}

	user := userInfo.(*models.User)
	return user, nil
}
