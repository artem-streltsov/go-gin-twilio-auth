package routes

import (
	"github.com/artem-streltsov/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, pc *controllers.ProfileController) {
	profile := r.Group("/profile")
	profile.Use(controllers.AuthMiddleware())
	{
		profile.GET("/", pc.GetUserProfile)
	}
}
