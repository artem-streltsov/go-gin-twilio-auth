package main

import (
	"os"

	"github.com/artem-streltsov/go-auth/controllers"
	"github.com/artem-streltsov/go-auth/database"
	"github.com/artem-streltsov/go-auth/models"
	"github.com/artem-streltsov/go-auth/routes"
	"github.com/artem-streltsov/go-auth/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.ConnectDB()
	db := database.DB
	db.AutoMigrate(&models.User{})

	twilioService := services.NewTwilioService()

	router := gin.Default()

	authController := controllers.NewAuthController(db, twilioService)
	routes.AuthRoutes(router, authController)

	profileController := controllers.NewProfileController(db)
	routes.ProfileRoutes(router, profileController)

	router.Run(os.Getenv("PORT"))
}
