package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/services"
)

func ListUsers(c *gin.Context) {
	users, err := services.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
