package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	RoomId      uint        `gorm:"type:int;not null" json:"room_id"`
	MettingRoom MettingRoom `gorm:"foreignKey:RoomId" json:"metting_rooms"`

	ResourceId uint     `gorm:"type:int;not null" json:"resource_id"`
	Resource   Resource `gorm:"foreignKey:ResourceId" json:"resource"`

	// Use the DATE data type in MySQL to store only the date portion.
	BookStartTime time.Time `gorm:"type:datetime;not null" json:"book_start_time"`
	BookEndTime   time.Time `gorm:"type:datetime;not null" json:"book_end_time"`

	Duration int `gorm:"type:int;not null" json:"duration"`
}
