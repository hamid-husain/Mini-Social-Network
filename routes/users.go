package routes

import (
	"github.com/gin-gonic/gin"

	"mini-social-network/controllers"
	"mini-social-network/middleware"
	"mini-social-network/services"
)

func APIRoutes(router *gin.Engine, service *services.Service) {

	ctrl := controllers.NewController(service)

	users := router.Group("/api")
	{
		users.POST("/create", ctrl.CreateUser)
		users.POST("/login", ctrl.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/logout", ctrl.Logout)
		protected.DELETE("/delete", ctrl.DeleteUser)
		protected.GET("/get_details", ctrl.GetUser)
		protected.GET("/list", ctrl.ListUsers)
		protected.PATCH("/update", ctrl.UpdateUser)
		protected.POST("/follow", ctrl.FollowUser)
		protected.POST("/unfollow", ctrl.UnfollowUser)
		protected.GET("/followings", ctrl.GetFollowing)
		protected.GET("/followers", ctrl.GetFollowers)
	}
}
