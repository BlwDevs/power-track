package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// UserParser representa um cliente parser associado a um usuário
type UserParserInverter struct {
	gorm.Model

	// Relacionamento com o usuário
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`

	// Status do cliente parser
	Active bool `json:"active" gorm:"type:boolean;default:true"`

	// Parametros do cliente parser (json)
	ParserParams json.RawMessage `json:"parser_params" gorm:"type:jsonb"`
	//FK ParserWorker

	ParserWorkerID uint         `json:"parser_worker_id" gorm:"not null"`
	ParserWorker   ParserWorker `gorm:"foreignKey:ParserWorkerID"`

	//FK Inverter
	InverterID uint     `json:"inverter_id" gorm:"not null"`
	Inverter   Inverter `gorm:"foreignKey:InverterID"`
}
