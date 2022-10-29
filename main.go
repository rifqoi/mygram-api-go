package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/routes"
	"github.com/jusidama18/mygram-api-go/config"
	"github.com/jusidama18/mygram-api-go/repository/gorm"
	"github.com/jusidama18/mygram-api-go/services"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	userRepo := gorm.NewUserRepository(db)
	userServices := services.NewUser(userRepo)
	userController := controllers.NewUser(userServices)

	router := gin.Default()
	app := routes.NewRouter(router, userController, userServices)
	app.Run()

}
