package router

import (
	"power-track/handler"
	"power-track/middleware"
	"power-track/service"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Inicializa middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.Cors())
	
	// Inicializa serviços e handlers
	inverterService := service.NewInverterService()
	inverterHandler := handler.NewInverterHandler(inverterService)

	stringpvService := service.NewStringpvService()
	stringpvHandler := handler.NewStringpvHandler(stringpvService)

	// Grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Rotas do inversor
		inverter := v1.Group("/inverters")
		{
			inverter.GET("/latest", inverterHandler.GetLastData)
			inverter.GET("/historical", inverterHandler.GetHistoricalData)
			inverter.GET("", inverterHandler.GetList)
			inverter.GET("/:id", inverterHandler.GetData)
		}

		// Rotas das strings fotovoltaicas
		stringpv := v1.Group("/strings")
		{
			stringpv.GET("/latest/:inverterId", stringpvHandler.GetLatest)
			stringpv.GET("/:inverterId", stringpvHandler.GetByInverter)
			stringpv.POST("", stringpvHandler.Create)
		}
	}
}
