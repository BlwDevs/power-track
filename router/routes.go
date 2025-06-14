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
	stringpvHandler := handler.NewStringpvHandler(stringpvService, parserWorkerRepo)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	UserParserInverterService := service.NewUserParserInverterService(UserParserInverterRepo)
	UserParserInverterHandler := handler.NewUserParserInverterHandler(UserParserInverterService, parserWorkerRepo)

	parserWorkerService := service.NewParserWorkerService(parserWorkerRepo)
	parserWorkerHandler := handler.NewParserWorkerHandler(parserWorkerService)

	// Grupo de rotas v1
	v1 := router.Group("/api/v1")
	{
		// Rotas do inversor
		inverter := v1.Group("/inverters")
		inverter.Use(middleware.AuthMiddleware(userService)) // Protege as rotas do inversor
		{
			inverter.POST("", inverterHandler.Create)           // Rota para criar um novo inversor
			inverter.GET("", inverterHandler.GetList)           // Rota para obter a lista de inversores
			inverter.GET("/:id", inverterHandler.GetData)       // Rota para obter dados de um inversor específico
			inverter.PUT("/:id", inverterHandler.Update)        // Rota para atualizar um inversor específico
			inverter.DELETE("/:id", inverterHandler.DeleteById) // Nova rota para deletar
		}

		// Rotas do ParserWorker
		parserWorker := v1.Group("/parser-worker")
		parserWorker.Use(middleware.AuthMiddleware(userService)) // Protege as rotas do parser worker
		{
			parserWorker.POST("", parserWorkerHandler.Create)     // Rota para criar um novo parser worker
			parserWorker.GET("", parserWorkerHandler.GetAll)      // Rota para obter todos os parser workers
			parserWorker.GET("/:id", parserWorkerHandler.GetByID) // Rota para obter um parser worker específico
			parserWorker.PUT("/:id", parserWorkerHandler.Update)  // Rota para atualizar um parser worker específico
			parserWorker.PUT("/api-key-refresh/:id", parserWorkerHandler.RefreshAPIKey)
			parserWorker.DELETE("/:id", parserWorkerHandler.Deactivate)                            // Rota para desativar um parser worker específico
			parserWorker.POST("/activate/:id", parserWorkerHandler.Activate)                       // Rota para ativar um parser worker específico
			parserWorker.GET("/manufacturer/:manufacturer", parserWorkerHandler.GetByManufacturer) // Rota para obter parser workers por fabricante

		}

		// Rotas das strings fotovoltaicas
		stringpv := v1.Group("/strings")
		stringpv.Use(middleware.AuthMiddleware(userService)) // Protege as rotas de strings fotovoltaicas
		{
			stringpv.GET("/latest/:inverterId", stringpvHandler.GetLatest)
			stringpv.GET("/historical/:inverterId", stringpvHandler.GetHistorical)
			stringpv.GET("/:inverterId", stringpvHandler.GetByInverter)
			stringpv.POST("", stringpvHandler.Create)
			// stringpv.POST("/csv-parser", stringpvHandler.CreateFromCSV)
		}

		// Rotas de usuários
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware(userService)) // Protege todas as rotas de usuários com autenticação
		{
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
			users.POST("/logout", userHandler.Logout)
		}

		// Rotas para login e autenticação
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/create", userHandler.Create)
			auth.GET("/growatt", UserParserInverterHandler.GetGrowattData) // Rota para obter dados do Growatt Parser Worker precisa de Api Key
			auth.POST("/batch", stringpvHandler.CreateMany)                // Rota para criar várias strings fotovoltaicas Parser Worker precisa de Api Key
		}

		// Rotas de clientes parser
		UserParserInverters := v1.Group("/user-parser")
		UserParserInverters.Use(middleware.AuthMiddleware(userService)) // Protege as rotas do UserParserInverter
		{
			UserParserInverters.POST("", UserParserInverterHandler.Create)
			UserParserInverters.GET("", UserParserInverterHandler.GetAll)
			UserParserInverters.GET("/:id", UserParserInverterHandler.GetByID)
			UserParserInverters.PUT("/:id", UserParserInverterHandler.Update)
			UserParserInverters.DELETE("/:id", UserParserInverterHandler.Delete)
			UserParserInverters.GET("/user/:userId", UserParserInverterHandler.GetByUserID)

		}
	}
}
