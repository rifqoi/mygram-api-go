package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
)

type Router struct {
	router *gin.Engine
	user   *controllers.UserController
}

func NewRouter(router *gin.Engine, user *controllers.UserController) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Run() {
	r.router.POST("/users/register", r.user.RegisterUser)

	r.router.Run(":8080")
}
