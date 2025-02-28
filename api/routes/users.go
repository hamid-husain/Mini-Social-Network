package routes

import (
	"project/api/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", controllers.ListUsers)
	}
}
