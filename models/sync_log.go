package models

import (
	"time"
	"gorm.io/gorm"
)

type SyncLog struct {
	gorm.Model
	TransactionID uint       `json:"transaction_id"`
	Status        string     `json:"status"`
	SyncedAt      *time.Time `json:"synced_at"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
}
