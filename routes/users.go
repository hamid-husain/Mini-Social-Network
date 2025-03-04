package routes

import (
	"github.com/gin-gonic/gin"

	"mini-social-network/controllers"
	"mini-social-network/middleware"
	"mini-social-network/services"
)

func APIRoutes(router *gin.Engine, service *services.Service) {
	users := router.Group("/api")
	{
		users.GET("/", controllers.ListUsers(service))
		users.POST("/create", controllers.CreateUser(service))
		users.POST("/login", controllers.Login(service))
	}

	protected := router.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/logout", controllers.Logout)
	}
}
