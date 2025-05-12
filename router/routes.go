package router

import (
	"power-track/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize the router
	v1 := router.Group("/api/v1")
	{
		// Define your routes here
		v1.GET("/lastData", handler.GetLastDataHandler)
		v1.GET("/historicalData", handler.GetHistoricalDataHandler)
		v1.GET("/listInversor", handler.GetListInversorHandler)
		v1.GET("/inversorData", handler.GetInversorDataHandler)
	}
	// r := gin.Default()
	// r.Use(middleware.Cors())
	// r.Use(middleware.Logger())
	// r.Use(middleware.Recovery())
	// r.Use(middleware.RequestId())
	// r.Use(middleware.Secure())
	// r.Use(middleware.SecureHeaders())
	// r.Use(middleware.SecureCookie())
	// r.Use(middleware.SecureSession())
	// r.Use(middleware.SecureCSRF())
	// r.Use(middleware.SecureXSS())
	// r.Use(middleware.SecureContentType())
	// r.Use(middleware.SecureContentSecurityPolicy())

	// Define your routes here
}
