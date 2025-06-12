package models

import (
	"gorm.io/gorm"
)

type Stringpv struct {
	gorm.Model
	Timestamp string  `json:"timestamp" gorm:"type:timestamp"`
	Index     int     `json:"index" gorm:"type:int"`
	StrNumber int     `json:"string_number" gorm:"type:int"`
	PPV       float64 `json:"ppv" gorm:"type:decimal(10,2)"`
	IPV       float64 `json:"ipv" gorm:"type:decimal(10,2)"`
	VPV       float64 `json:"vpv" gorm:"type:decimal(10,2)"`

	InverterID uint     `json:"inverter_id" gorm:"type:int;not null"`
	Inverter   Inverter `gorm:"foreignKey:InverterID" json:"-"`
}
