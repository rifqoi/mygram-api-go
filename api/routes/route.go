package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/middlewares"
	"github.com/jusidama18/mygram-api-go/services"
)

type Router struct {
	router  *gin.Engine
	userCtl *controllers.UserController
	userSvc *services.UserService
}

func NewRouter(router *gin.Engine, user *controllers.UserController, userSvc *services.UserService) *Router {
	return &Router{
		router:  router,
		userCtl: user,
		userSvc: userSvc,
	}
}

func (r *Router) Run() {
	r.router.POST("/users/register", r.userCtl.RegisterUser)
	r.router.POST("/users/login", r.userCtl.Login)

	// Check middleware
	r.router.GET("/check", middlewares.Authorization(r.userSvc), r.userCtl.Check)

	r.router.Run(":8080")
}
