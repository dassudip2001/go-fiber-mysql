package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	RoomId uint        `gorm:"type:int;not null" json:"room_id"`
	Room   MettingRoom `gorm:"foreignKey:RoomId" json:"room"`

	ResourceId uint     `gorm:"type:int;not null" json:"resource_id"`
	Resource   Resource `gorm:"foreignKey:ResourceId" json:"resource"`

	BookStartTime time.Time `gorm:"type:datetime;not null" json:"book_start_time"`
	BookEndTime   time.Time `gorm:"type:datetime;not null" json:"book_end_time"`

	Duration int `gorm:"type:int;not null" json:"duration"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
