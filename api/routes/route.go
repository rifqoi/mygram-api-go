package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/middlewares"
)

type Router struct {
	router     *gin.Engine
	user       *controllers.UserController
	photo      *controllers.PhotoController
	middleware *middlewares.Middleware
}

func NewRouter(router *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController, middleware *middlewares.Middleware) *Router {
	return &Router{
		router:     router,
		user:       user,
		photo:      photo,
		middleware: middleware,
	}
}

func (r *Router) Run() {
	userRoutes := r.router.Group("/users")
	userRoutes.POST("/register", r.user.RegisterUser)
	userRoutes.POST("/login", r.user.Login)
	userRoutes.PUT("/", r.middleware.Authorization, r.user.UpdateUser)
	userRoutes.DELETE("/", r.middleware.Authorization, r.user.DeleteUser)

	photoRoutes := r.router.Group("/photos").Use(r.middleware.Authorization)
	photoRoutes.POST("/", r.photo.CreatePhoto)
	photoRoutes.GET("/", r.photo.GetAllPhotos)
	photoRoutes.PUT("/:id", r.photo.UpdatePhoto)
	photoRoutes.DELETE("/:id", r.photo.DeletePhoto)

	// Check middleware
	r.router.GET("/check", r.middleware.Authorization, r.user.Check)

	r.router.Run(":8080")
}
