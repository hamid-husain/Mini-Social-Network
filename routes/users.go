package routes

import (
	"github.com/gin-gonic/gin"

	"mini-social-network/controllers"
	"mini-social-network/middleware"
	"mini-social-network/services"
)

func APIRoutes(router *gin.Engine, authService *services.AuthService, userService *services.UserService) {

	authCtrl := controllers.NewAuthController(authService)
	userCtrl := controllers.NewUserController(userService)

	users := router.Group("/api/v1")
	{
		users.POST("/create", authCtrl.CreateUser)
		users.POST("/login", authCtrl.Login)
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/logout", authCtrl.Logout)
		protected.DELETE("/delete", userCtrl.DeleteUser)
		protected.GET("/get_details", userCtrl.GetUser)
		protected.GET("/list", userCtrl.ListUsers)
		protected.PATCH("/update", userCtrl.UpdateUser)
		protected.POST("/follow", userCtrl.FollowUser)
		protected.POST("/unfollow", userCtrl.UnfollowUser)
		protected.GET("/followings", userCtrl.GetFollowing)
		protected.GET("/followers", userCtrl.GetFollowers)
		protected.POST("/update_password", userCtrl.UpdatePassword)
	}
}
