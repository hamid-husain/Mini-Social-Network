package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/config"
	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
)

type Controller struct {
	Service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{Service: service}
}

func (ctrl *Controller) CreateUser(c *gin.Context) {
	var req serializers.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
		return
	}

	response, err := ctrl.Service.CreateUserWithDetails(&req)
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

func (ctrl *Controller) Login(c *gin.Context) {
	var req serializers.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ParseValidationErrors(err))
		return
	}

	user, err := ctrl.Service.VerifyUserCredentials(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToGenerateToken})
		return
	}

	userResponse := serializers.SerializeUserLoginResponse(*user)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	response := serializers.SerializeLoginResponse(userResponse, tokenResponse)
	c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessLogOut})
}
