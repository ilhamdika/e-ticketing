package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
	User    User    `gorm:"foreignKey:UserID"`
}
