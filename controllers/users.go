package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
)

func ListUsers(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.ListUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetUser(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		response, err := service.GetUserByID(userID.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func DeleteUser(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		response, err := service.DeleteUserByID(userID.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, response)
	}
}

func UpdateUser(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}
		var req serializers.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrorsMap := utils.ParseValidationErrors(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
			return
		}

		response, err := service.UpdateUserByID(userID.(uint), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)

	}
}
