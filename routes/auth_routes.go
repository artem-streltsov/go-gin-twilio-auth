package routes

import (
	"github.com/artem-streltsov/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, ac *controllers.AuthController) {
	r.POST("/register", ac.RegisterUser)
	r.POST("/login", ac.LoginUser)
	r.POST("/verify", ac.VerifyAndGenerateJWT)
}
