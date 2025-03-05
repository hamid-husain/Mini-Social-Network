package main

import (
	"fmt"
	"log"
	"mini-social-network/config"
	"mini-social-network/db"
	"mini-social-network/routes"
	"mini-social-network/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.LoadConfig()

	db.ConnectDatabase()

	validator := validator.New()
	utils.RegisterCustomValidators(validator)

	router := gin.Default()

	routes.UsersRoutes(router)

	if err := router.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		fmt.Println("Server is running on port " + config.AppConfig.ServerPort)
	}
}
