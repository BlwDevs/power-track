package models

import (
	"gorm.io/gorm"
)

type Status string

const (
	Online  Status = "online"
	Offline Status = "offline"
	Unknown Status = "unknown"
)

type Inverter struct {
	gorm.Model

	SerialNumber string  `json:"serial_number" gorm:"type:varchar(50)"`
	Status       Status  `json:"status" gorm:"type:varchar(50)"`
	Power        float64 `json:"power" gorm:"type:decimal(10,2)"`
	TotalEnergy  float64 `json:"total_energy" gorm:"type:decimal(10,2)"`
}