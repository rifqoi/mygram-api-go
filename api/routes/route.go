package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/middlewares"
)

type Router struct {
	router     *gin.Engine
	user       *controllers.UserController
	middleware *middlewares.Middleware
}

func NewRouter(router *gin.Engine, user *controllers.UserController, middleware *middlewares.Middleware) *Router {
	return &Router{
		router:     router,
		user:       user,
		middleware: middleware,
	}
}

func (r *Router) Run() {
	r.router.POST("/users/register", r.user.RegisterUser)
	r.router.POST("/users/login", r.user.Login)

	// Check middleware
	r.router.GET("/check", r.middleware.Authorization(), r.user.Check)
	r.router.PUT("/users", r.middleware.Authorization(), r.user.UpdateUser)
	r.router.DELETE("/users", r.middleware.Authorization(), r.user.DeleteUser)

	r.router.Run(":8080")
}
