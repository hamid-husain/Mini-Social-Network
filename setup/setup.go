package setup

import (
	"fmt"
	"log"
	"mini-social-network/config"
	"mini-social-network/db"
	"mini-social-network/middleware"
	"mini-social-network/routes"
	"mini-social-network/services"
	"mini-social-network/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	return router
}

func init() {

	config.LoadConfig()

	db.ConnectDatabase()

	authService := services.NewAuthService(db.DB)
	userService := services.NewUserService(db.DB)

	router := SetupRouter()
	routes.APIRoutes(router, authService, userService)

	validator := validator.New()
	utils.RegisterCustomValidators(validator)

	if err := router.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		fmt.Println("Server is running on port " + config.AppConfig.ServerPort)
	}
}
