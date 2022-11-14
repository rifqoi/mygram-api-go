package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/services"
)

type CommentController struct {
	svc *services.CommentService
}

func NewCommentController(svc *services.CommentService) *CommentController {
	return &CommentController{
		svc: svc,
	}
}

func (cm *CommentController) CreateComment(c *gin.Context) {
	var req parameters.CreateComment

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(&req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	resComment, err := cm.svc.CreateComment(req, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusCreated, resComment, "comment successfully created")
}

func (cm *CommentController) GetAllComment(c *gin.Context) {
	resComment, err := cm.svc.GetAllComment()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.SuccessWithData(c, http.StatusOK, resComment, "successfully get all comments")
}

func (cm *CommentController) UpdateComment(c *gin.Context) {
	var req parameters.UpdateComment

	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	updatedComment, err := cm.svc.UpdateComment(req, commentID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusAccepted, updatedComment, "comment updated successfully.")
}

func (cm *CommentController) DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = cm.svc.DeleteComment(commentID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusAccepted, "comment successfully deleted.")
}
