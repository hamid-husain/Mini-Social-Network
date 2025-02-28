package main

import (
	"fmt"
	"log"
	"project/api/routes"
	"project/config"
	"project/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db.ConnectDatabase()

	router := gin.Default()

	routes.UsersRoutes(router)

	if err := router.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		fmt.Println("Server is running on port " + config.AppConfig.ServerPort)
	}
}
