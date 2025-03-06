package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"mini-social-network/constants"
	"mini-social-network/serializers"
	"mini-social-network/utils"
)

func (ctrl *Controller) ListUsers(c *gin.Context) {
	response, err := ctrl.Service.ListUsers()
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) GetUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	response, err := ctrl.Service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) DeleteUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
		return
	}

	response, err := ctrl.Service.DeleteUserByID(userID.(uint))
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

func (ctrl *Controller) UpdateUser(c *gin.Context) {
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

	response, err := ctrl.Service.UpdateUserByID(userID.(uint), &req)
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

func (ctrl *Controller) FollowUser(c *gin.Context) {
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

	err := ctrl.Service.FollowUsersByID(userID.(uint), req.UserIDs)
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

func (ctrl *Controller) UnfollowUser(c *gin.Context) {
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

	err := ctrl.Service.UnfollowUsersByID(userID.(uint), req.UserIDs)
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
