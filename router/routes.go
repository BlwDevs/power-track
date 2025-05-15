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
	userRepo := repository.NewUserRepository(db)
	parserClientRepo := repository.NewParserClientRepository(db)
	
	// Inicializa serviços e handlers
	inverterService := service.NewInverterService(inverterRepo)
	inverterHandler := handler.NewInverterHandler(inverterService)

	stringpvService := service.NewStringpvService(stringpvRepo)
	stringpvHandler := handler.NewStringpvHandler(stringpvService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	parserClientService := service.NewParserClientService(parserClientRepo)
	parserClientHandler := handler.NewParserClientHandler(parserClientService)

	// Grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Rotas do inversor
		inverter := v1.Group("/inverters")
		{
			inverter.POST("", inverterHandler.Create)
			inverter.GET("/latest", inverterHandler.GetLastData)
			inverter.GET("/historical", inverterHandler.GetHistoricalData)
			inverter.GET("", inverterHandler.GetList)
			inverter.GET("/:id", inverterHandler.GetData)
			inverter.DELETE("/:id", inverterHandler.DeleteById)  // Nova rota para deletar
		}

		// Rotas das strings fotovoltaicas
		stringpv := v1.Group("/strings")
		{
			stringpv.GET("/latest/:inverterId", stringpvHandler.GetLatest)
			stringpv.GET("/:inverterId", stringpvHandler.GetByInverter)
			stringpv.POST("", stringpvHandler.Create)
			// stringpv.POST("/csv-parser", stringpvHandler.CreateFromCSV)
		}

		// Rotas de usuários
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}

		// Rotas de clientes parser
		parserClients := v1.Group("/parser-clients")
		{
			parserClients.POST("", parserClientHandler.Create)
			parserClients.GET("", parserClientHandler.GetAll)
			parserClients.GET("/:id", parserClientHandler.GetByID)
			parserClients.PUT("/:id", parserClientHandler.Update)
			parserClients.DELETE("/:id", parserClientHandler.Delete)
			parserClients.GET("/user/:userId", parserClientHandler.GetByUserID)
		}
	}
}
