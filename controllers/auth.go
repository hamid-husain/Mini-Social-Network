package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

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

func Login(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req serializers.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, utils.ParseValidationErrors(err))
			return
		}

		user, err := service.VerifyUserCredentials(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials})
			return
		}

		token, expiryTime, err := utils.GenerateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToGenerateToken})
			return
		}

		bearerToken := "Bearer " + token

		userResponse := serializers.SerializeUserLoginResponse(*user)
		tokenResponse := serializers.SerializeToken(bearerToken, expiryTime)
		response := serializers.SerializeLoginResponse(userResponse, tokenResponse)
		c.JSON(http.StatusOK, response)
	}
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessLogOut})
}

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
