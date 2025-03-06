package setup

import (
	"fmt"
	"log"
	"mini-social-network/config"
	"mini-social-network/db"
	"mini-social-network/db/base_model"
	"mini-social-network/middleware"
	"mini-social-network/routes"
	"mini-social-network/services"
	"mini-social-network/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	config.LoadConfig()

	db.ConnectDatabase()

	db.DB.AutoMigrate(&base_model.User{}, &base_model.OfficeDetail{}, &base_model.ResidentialDetail{}, &base_model.UserFollowing{})

	service := services.NewService(db.DB)

	validator := validator.New()
	utils.RegisterCustomValidators(validator)

	router := gin.Default()

	routes.APIRoutes(router, service)
	router.Use(middleware.CORSMiddleware())

	if err := router.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		fmt.Println("Server is running on port " + config.AppConfig.ServerPort)
	}
}
