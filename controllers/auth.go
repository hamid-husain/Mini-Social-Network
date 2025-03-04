package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"time"

	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
)

func CreateUser(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req serializers.SignUpRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrorsMap := utils.ParseValidationErrors(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
			return
		}

		response, err := service.CreateUserWithDetails(&req)
		if err != nil {
			if err.Error() == constants.ErrEmailAlreadyExists {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusCreated, response)
	}
}

func Login(c *gin.Context) {
	var req serializers.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ParseValidationErrors(err))
		return
	}

	user, err := services.VerifyUserCredentials(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToGenerateToken})
		return
	}

	c.SetCookie("token", token, int(expiryTime-time.Now().Unix()), "/", "localhost", false, true)

	userResponse := serializers.SerializeLoginResponse(*user)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	c.JSON(http.StatusOK, gin.H{
		"user_details": userResponse,
		"token":        tokenResponse,
	})
}
