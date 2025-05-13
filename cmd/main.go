package main

import (
	"power-track/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Initialize DB connection
	db, err := initDB() // You need to implement this function
	if err != nil {
		panic("Failed to connect to database")
	}

	// Initialize Gin router
	r := gin.Default()

	// Initialize routes with both router and db
	router.InitializeRoutes(r, db)

	// Start server
	r.Run(":8080")
}

// initDB initializes and returns a database connection
func initDB() (*gorm.DB, error) {
	// Implement your database connection logic here
	// Example:
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	// TODO: Implement your actual database connection
	return nil, nil
}