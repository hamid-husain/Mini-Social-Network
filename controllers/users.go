package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/constants"
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

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	_, err := services.GetUserByID(userID)
	if err != nil {
		if err.Error() == constants.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	response, err := services.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
