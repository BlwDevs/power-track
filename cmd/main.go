package main

import (
	"power-track/database"
	"power-track/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB connection
	db := database.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Initialize routes with both router and db
	router.InitializeRoutes(r, db)

	// Start server
	r.Run(":8080")
}