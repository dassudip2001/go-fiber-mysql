package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"size:255;not null;unique" json:"name"`
	Description string   `gorm:"size:255; default:null" json:"description"`
	Price       int      `gorm:"not null" json:"price"`
	Stock       int      `gorm:"not null" json:"stock"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
