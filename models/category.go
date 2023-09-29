package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255; default:null" json:"description"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
