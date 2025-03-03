package routes

import (
	"mini-social-network/controllers"
	"mini-social-network/services"

	"github.com/gin-gonic/gin"
)

func APIRoutes(router *gin.Engine, service *services.Service) {
	users := router.Group("/api")
	{
		users.GET("/", controllers.ListUsers(service))
		users.POST("/create", controllers.CreateUser(service))
		users.POST("/login", controllers.Login)
	}
}
