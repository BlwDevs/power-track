package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"power-track/models"
)

func Connect() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Conectado ao banco com sucesso!")

	// AutoMigrate models
	err = db.AutoMigrate(&models.Inverter{}, &models.Stringpv{}, &models.User{}, &models.UserParserInverter{}, &models.ParserWorker{})
	if err != nil {
		log.Fatal("Erro ao fazer AutoMigrate:", err)
	}

	return db
}
