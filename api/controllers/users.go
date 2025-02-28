package controllers

import (
	"net/http"
	"project/db"
	"project/internal/models/users"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	var users []users.User
	db.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
