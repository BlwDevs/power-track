package handler

import "github.com/gin-gonic/gin"

func GetLastDataHandler(ctx *gin.Context) {
	// Handle the request and send a response
	ctx.JSON(200, gin.H{
		"message": "Get Last Data",
	})
}

func GetHistoricalDataHandler(ctx *gin.Context) {
	// Handle the request and send a response
	ctx.JSON(200, gin.H{
		"message": "Get Historical Data",
	})
}

func GetListInversorHandler(ctx *gin.Context) {
	// Handle the request and send a response
	ctx.JSON(200, gin.H{
		"message": "Get Inversor List",
	})
}

func GetInversorDataHandler(ctx *gin.Context) {
	// Handle the request and send a response
	ctx.JSON(200, gin.H{
		"message": "Get Inversor Data",
	})
}
