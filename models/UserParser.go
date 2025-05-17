package models

import (
	"gorm.io/gorm"
)

// UserParser representa um cliente parser associado a um usuário
type UserParser struct {
	gorm.Model

	// Relacionamento com o usuário
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`

	// Status do cliente parser
	Active bool `json:"active" gorm:"type:boolean;default:true"`

	// Parametros do cliente parser (json)
	ParserParams string `json:"parser_params" gorm:"type:jsonb"`

	//FK ParserWorker

	ParserWorkerID uint         `json:"parser_worker_id" gorm:"not null"`
	ParserWorker   ParserWorker `gorm:"foreignKey:ParserWorkerID"`
}
