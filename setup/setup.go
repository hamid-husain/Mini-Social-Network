package setup

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"log"

	"mini-social-network/config"
	"mini-social-network/db"
	"mini-social-network/db/base_model"
	"mini-social-network/middleware"
	"mini-social-network/routes"
	"mini-social-network/services"
	"mini-social-network/validator"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	return router
}

func init() {

	config.LoadConfig()

	db.ConnectDatabase()
	db.DB.AutoMigrate(&base_model.User{}, &base_model.ResidentialDetail{}, &base_model.OfficeDetail{}, &base_model.UserFollowing{})

	authService := services.NewAuthService(db.DB)
	userService := services.NewUserService(db.DB)

	router := SetupRouter()
	routes.APIRoutes(router, authService, userService)

	validator.RegisterCustomValidators()

	if err := router.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		fmt.Println("Server is running on port " + config.AppConfig.ServerPort)
	}
}
