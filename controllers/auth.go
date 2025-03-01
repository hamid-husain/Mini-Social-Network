package controllers

import (
	"mini-social-network/models/baseModel"
	"mini-social-network/models/users"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req serializers.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := users.User{
		User: baseModel.User{
			Email:         req.Email,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			DateOfBirth:   req.DateOfBirth,
			Gender:        req.Gender,
			MaritalStatus: req.MaritalStatus,
			Password:      req.Password,
		},
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	userResponse := serializers.SerializeUser(user)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": tokenResponse, "user": userResponse})
}
