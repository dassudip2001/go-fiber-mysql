package models

import (
	"time"

	"gorm.io/gorm"
)

type MettingRoom struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(256);not null" json:"name"`
	Capacity      int     `gorm:"type:int;not null" json:"capacity"`
	Configuration string  `gorm:"type:varchar(100);not null" json:"configuration"`
	IsAvailable   bool    `gorm:"type:tinyint(1);not null" json:"is_available"`
	Description   *string `gorm:"type:varchar(255)" json:"description"`

	LocationId uint     `gorm:"type:int;not null" json:"location_id"`
	Location   Location `gorm:"foreignKey:LocationId" json:"location"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
