package models

import (
	"gorm.io/gorm"
)

// ParserWorker representa um programa para processar dados de strings fotovoltaicas e ajustar ao formato do banco de dados
type ParserWorker struct {
	gorm.Model

	//Fabricante do inversor
	Manufacturer string `json:"manufacturer" gorm:"type:varchar(100);not null"`
	// Chave da API para autenticação
	ApiKey string `json:"api_key" gorm:"type:varchar(255);not null;uniqueIndex"`
}
