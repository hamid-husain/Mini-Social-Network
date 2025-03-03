package controllers

import (
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req serializers.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ParseValidationErrors(err))
		return
	}

	if err := utils.ValidateDOB(req.UserDetails.DateOfBirth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"date_of_birth": err.Error()}})
		return
	}

	user, office, resident, err := services.CreateUserWithDetails(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	userResponse := serializers.SerializeResponse(*user, *resident, *office)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	c.JSON(http.StatusCreated, gin.H{
		"id":            user.ID,
		"user_id":       user.ID,
		"email":         user.Email,
		"last_modified": user.UpdatedAt,
		"user_details":  userResponse,
		"token":         tokenResponse,
	})
}
