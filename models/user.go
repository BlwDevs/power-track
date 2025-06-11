package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Domínio para o atributo plano do usuário
type Plan int

// Variáveis numéricas para o plano do usuário
const (
	Free        Plan = 0
	OneInverter Plan = 1
	TwoInverter Plan = 2
)

type Role string

const (
	Admin  Role = "admin"
	Common Role = "common"
)

type User struct {
	gorm.Model
	CPF      string `json:"cpf" gorm:"type:varchar(14);uniqueIndex"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Address  string `json:"address" gorm:"type:varchar(255)"`
	Plan     Plan   `json:"plan" gorm:"type:int;default:0"`
	Phone    string `json:"phone" gorm:"type:varchar(20);uniqueIndex"`
	Role     Role   `json:"role" gorm:"type:varchar(10);default:'common';check:role IN ('admin','common')"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);uniqueIndex;default:null;null"`
}

// HashPassword cria um hash da senha do usuário
func (u *User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
