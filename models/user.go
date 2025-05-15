package models

import (
	"gorm.io/gorm"
)

//Domínio para o atributo plano do usuário
type Plan int
//Variáveis numéricas para o plano do usuário
const (
	Free  Plan = 0
	OneInverter Plan = 1
	TwoInverter Plan = 2
)
type User struct {
	gorm.Model
	CPF     string `json:"cpf" gorm:"type:varchar(14);uniqueIndex"`
	Name    string `json:"name" gorm:"type:varchar(100)"`
	Email   string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Address string `json:"address" gorm:"type:varchar(255)"`
	Plan    Plan   `json:"plan" gorm:"type:int;default:0"`
	Phone   string `json:"phone" gorm:"type:varchar(20);uniqueIndex"`
} 