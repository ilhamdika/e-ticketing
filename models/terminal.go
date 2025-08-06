package models

import "gorm.io/gorm"

type Terminal struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	CreatedBy uint  `json:"created_by"`
}
