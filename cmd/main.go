package main

import (
	"power-track/database"
	"power-track/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	Server := gin.Default()

	// Initialize the router
	router.InitializeRoutes(Server)
	
	Server.Run(":8080")
}