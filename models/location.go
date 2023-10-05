package models

import (
	"time"

	"gorm.io/gorm"
)




type Location struct {
    gorm.Model
    Name       string
    ParentID   *uint
    Children   []Location `gorm:"foreignkey:ParentID"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}

// type Location struct {
//     gorm.Model
//     Name       string
//     ParentID   *uint
//     Children   []Location `gorm:"foreignkey:ParentID"`
//     CreatedAt  time.Time
//     UpdatedAt  time.Time
// }

// // Add a foreign key constraint for the ParentID field
// func (l *Location) BeforeSave(tx *gorm.DB) error {
//     if l.ParentID != nil && *l.ParentID == l.ID {
//         return errors.New("cannot set ParentID to the same as ID")
//     }
//     return nil
// }
