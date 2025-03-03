package controllers

import (
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req serializers.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ParseValidationErrors(err))
		return
	}

	user, err := services.VerifyUserCredentials(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	userResponse := serializers.SerializeLoginResponse(*user)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	c.JSON(http.StatusOK, gin.H{
		"user_details": userResponse,
		"token":        tokenResponse,
	})
}
