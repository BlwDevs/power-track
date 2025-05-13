package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"power-track/models"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=goapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Conectado ao banco com sucesso!")

	// AutoMigrate models
	err = db.AutoMigrate(&models.Inverter{}, &models.Stringpv{})
	if err != nil {
		log.Fatal("Erro ao fazer AutoMigrate:", err)
	}

	return db
}