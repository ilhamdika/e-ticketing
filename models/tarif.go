package models

import "gorm.io/gorm"

type Tarif struct {
	gorm.Model
	OriginTerminalID uint     `json:"origin_terminal_id"`
	DestTerminalID   uint     `json:"dest_terminal_id"`
	Price            float64  `json:"price"`
	
	OriginTerminal Terminal `gorm:"foreignKey:OriginTerminalID"`
	DestTerminal   Terminal `gorm:"foreignKey:DestTerminalID"`
}
