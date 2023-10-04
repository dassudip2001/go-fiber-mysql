package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description *string

	CreatedAt time.Time
	UpdatedAt time.Time
}
