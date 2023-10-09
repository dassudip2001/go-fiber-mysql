package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name        string `gorm:"type:varchar(256);not null" json:"name"`
	IsAvailable bool   `gorm:"default:true" json:"is_available"`
}
