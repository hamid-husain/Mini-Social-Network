package routes

import (
	"mini-social-network/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	users := router.Group("/api")
	{
		users.GET("/", controllers.ListUsers)
		users.POST("/create", controllers.CreateUser)
	}
}
