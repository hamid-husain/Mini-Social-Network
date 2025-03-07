package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (ctrl *UserController) ListUsers(c *gin.Context) {
	response, err := ctrl.UserService.ListUsers()
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	response, err := ctrl.UserService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	response, err := ctrl.UserService.DeleteUserByID(userID.(uint))
	if err != nil {
		if err.Status() == http.StatusNotFound {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServerError})
		return

	}

	c.SetCookie("token", "", -1, "/", constants.Domain, false, true)
	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) UpdatePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	var req serializers.PasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrorsMap})
		return
	}

	err := ctrl.UserService.UpdatePasswordByID(userID.(uint), req)

	if err != nil {
		if err.Status() == http.StatusInternalServerError {
			c.JSON(err.Status(), gin.H{"error": constants.ErrInternalServerError})
		}
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessPasswordUpdated})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
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

	response, err := ctrl.UserService.UpdateUserByID(userID.(uint), &req)
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

func (ctrl *UserController) FollowUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	var req serializers.UserFollowRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrorsMap})
		return
	}

	err := ctrl.UserService.FollowUsersByID(userID.(uint), req.UserIDs)
	if err != nil {
		if err.Status() == http.StatusBadRequest {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessUserFollowed})
}

func (ctrl *UserController) UnfollowUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	var req serializers.UserFollowRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrorsMap := utils.ParseValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrorsMap})
		return
	}

	err := ctrl.UserService.UnfollowUsersByID(userID.(uint), req.UserIDs)
	if err != nil {
		if err.Status() == http.StatusBadRequest {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": constants.SuccessUserUnfollowed})
}

func (ctrl *UserController) GetFollowers(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	followers, err := ctrl.UserService.GetFollowersByID(userID.(uint))
	if err != nil {
		if err.Status() == http.StatusNotFound {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_ids": followers})
}

func (ctrl *UserController) GetFollowing(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}
	following, err := ctrl.UserService.GetFollowing(userID.(uint))
	if err != nil {
		if err.Status() == http.StatusNotFound {
			c.JSON(err.Status(), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_ids": following})
}
