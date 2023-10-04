package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description *string
	Price       int
	Stock       int
	CategoryID  uint
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
