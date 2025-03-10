package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

type AuthControllerInt interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

func (ctrl *AuthController) CreateUser(c *gin.Context) {
	var req serializers.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
		return
	}

	response, err := ctrl.AuthService.CreateUserWithDetails(&req)
	if err != nil {
		if err.Error() == constants.ErrEmailAlreadyExists {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{"email": []string{err.Error()}}})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req serializers.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ParseValidationErrors(err))
		return
	}

	response, err := ctrl.AuthService.LoginHandler(req)
	if err != nil {
		if err.Status() == http.StatusInternalServerError {
			c.JSON(err.Status(), gin.H{"error": constants.ErrInternalServerError})
			return
		}

		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessLogOut})
}
