package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Name         string   `json:"name"`
	PhoneNumber  string   `json:"phone_number"`
	Email        string   `json:"email"`
	Cards        []Card   `gorm:"foreignKey:UserID"`
}
