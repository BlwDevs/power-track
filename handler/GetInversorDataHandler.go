package handler

import "github.com/gin-gonic/gin"

func GetInversorDataHandler(ctx *gin.Context) {
	// Handle the request and send a response
	ctx.JSON(200, gin.H{
		"message": "Get Inversor Data",
	})
}
