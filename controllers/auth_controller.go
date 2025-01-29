package controllers

import (
	"net/http"

	"github.com/artem-streltsov/go-auth/models"
	"github.com/artem-streltsov/go-auth/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB     *gorm.DB
	Twilio *services.TwilioService
}

func NewAuthController(db *gorm.DB, twilio *services.TwilioService) *AuthController {
	return &AuthController{DB: db, Twilio: twilio}
}

type RegisterInput struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:    input.Name,
		Surname: input.Surname,
		Phone:   input.Phone,
	}

	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type PhoneInput struct {
	Phone string `json:"phone" binding:"required"`
}

func (ac *AuthController) StartPhoneVerification(c *gin.Context) {
	var input PhoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.Twilio.StartVerification(input.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
}

type VerifyCodeInput struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func (ac *AuthController) VerifyPhone(c *gin.Context) {
	var input VerifyCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := ac.Twilio.CheckVerification(input.Phone, input.Code)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	var user models.User
	if err := ac.DB.Model(&user).Where("phone = ?", input.Phone).Update("verified", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone number verified successfully"})
}
