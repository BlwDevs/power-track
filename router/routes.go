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
	UserParserInverterRepo := repository.NewUserParserInverterRepository(db)
	parserWorkerRepo := repository.NewParserWorkerRepository(db)

	// Inicializa serviços e handlers
	inverterService := service.NewInverterService(inverterRepo)
	inverterHandler := handler.NewInverterHandler(inverterService)

	stringpvService := service.NewStringpvService(stringpvRepo)
	stringpvHandler := handler.NewStringpvHandler(stringpvService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	UserParserInverterService := service.NewUserParserInverterService(UserParserInverterRepo)
	UserParserInverterHandler := handler.NewUserParserInverterHandler(UserParserInverterService)

	parserWorkerService := service.NewParserWorkerService(parserWorkerRepo)
	parserWorkerHandler := handler.NewParserWorkerHandler(parserWorkerService)

	// Grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Rotas do inversor
		inverter := v1.Group("/inverters")
		{
			inverter.POST("", inverterHandler.Create)
			inverter.GET("", inverterHandler.GetList)
			inverter.GET("/:id", inverterHandler.GetData)
			inverter.PUT("/:id", inverterHandler.Update)
			inverter.DELETE("/:id", inverterHandler.DeleteById) // Nova rota para deletar
		}

		// Rotas do ParserWorker
		parserWorker := v1.Group("/parser-worker")
		{
			parserWorker.POST("", parserWorkerHandler.Create)
			parserWorker.GET("", parserWorkerHandler.GetAll)
			parserWorker.GET("/:id", parserWorkerHandler.GetByID)
			parserWorker.PUT("/:id", parserWorkerHandler.Update)
			parserWorker.DELETE("/:id", parserWorkerHandler.Deactivate)
			parserWorker.POST("/activate/:id", parserWorkerHandler.Activate)
			parserWorker.GET("/manufacturer/:manufacturer", parserWorkerHandler.GetByManufacturer)

		}

		// Rotas das strings fotovoltaicas
		stringpv := v1.Group("/strings")
		{
			stringpv.GET("/latest/:inverterId", stringpvHandler.GetLatest)
			stringpv.GET("/historical/:inverterId", stringpvHandler.GetHistorical)
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
		UserParserInverters := v1.Group("/user-parser")
		{
			UserParserInverters.POST("", UserParserInverterHandler.Create)
			UserParserInverters.GET("", UserParserInverterHandler.GetAll)
			UserParserInverters.GET("/:id", UserParserInverterHandler.GetByID)
			UserParserInverters.PUT("/:id", UserParserInverterHandler.Update)
			UserParserInverters.DELETE("/:id", UserParserInverterHandler.Delete)
			UserParserInverters.GET("/user/:userId", UserParserInverterHandler.GetByUserID)
			parserWorker.GET("/growatt", UserParserInverterHandler.GetGrowattData)
		}
	}
}
