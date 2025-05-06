package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Initialize the router
	v1 := router.Group("/api/v1");
	{
		// Define your routes here
		 v1.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World! v1",
			})
			
		 })
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