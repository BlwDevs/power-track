package models

import (
	"gorm.io/gorm"
)

// ParserClient representa um cliente com chave de parser
type ParserClient struct {
	gorm.Model

	// ParserKey é a chave de análise/parser
	ParserKey string `json:"parser_key" gorm:"type:varchar(255);not null;uniqueIndex"`
	
	// Relacionamento com o usuário
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
	
	// Status do cliente parser
	Active bool `json:"active" gorm:"type:boolean;default:true"`
	
	// Descrição ou nome do cliente
	Name string `json:"name" gorm:"type:varchar(100)"`
} 