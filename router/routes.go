package router

import (
	"power-track/handler"
	"power-track/middleware"
	"power-track/repository"
	"power-track/service"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	// Inicializa middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.Cors())
	
	// Inicializa repositories
	inverterRepo := repository.NewInverterRepository(db)
	stringpvRepo := repository.NewStringpvRepository(db)
	
	// Inicializa serviços e handlers
	inverterService := service.NewInverterService(inverterRepo)
	inverterHandler := handler.NewInverterHandler(inverterService)

	stringpvService := service.NewStringpvService(stringpvRepo)
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
			// stringpv.POST("/csv-parser", stringpvHandler.CreateFromCSV)
		}
	}
}
