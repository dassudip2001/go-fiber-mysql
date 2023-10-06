package models

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name        string `gorm:"type:varchar(256);not null" json:"name"`
	IsAvailable bool   `gorm:"default:true" json:"is_available"`
	CreateAt    time.Time
	UpdateAt    time.Time
}
