package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/services"
)

func ListUsers(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := service.FindUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}
