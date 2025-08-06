package models

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CardID            uint       `json:"card_id"`
	CheckinTerminalID uint       `json:"checkin_terminal_id"`
	CheckoutTerminalID uint      `json:"checkout_terminal_id"`
	CheckinTime       *time.Time `json:"checkin_time"`
	CheckoutTime      *time.Time `json:"checkout_time"`
	FareApplied       float64    `json:"fare_applied"`
	Synced            bool       `json:"synced"`

	Card             Card     `gorm:"foreignKey:CardID"`
	CheckinTerminal  Terminal `gorm:"foreignKey:CheckinTerminalID"`
	CheckoutTerminal Terminal `gorm:"foreignKey:CheckoutTerminalID"`
}
