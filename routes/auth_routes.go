package routes

import (
	"github.com/artem-streltsov/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, ac *controllers.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", ac.Register)
		auth.POST("/send-verification", ac.StartPhoneVerification)
		auth.POST("/verify", ac.VerifyPhone)
	}
}
