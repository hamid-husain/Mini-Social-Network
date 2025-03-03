package controllers

import (
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(c *gin.Context) {
	validator := validator.New()
	utils.RegisterCustomValidators(validator)

	var req serializers.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
		return
	}

	response, err := services.CreateUserWithDetails(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
