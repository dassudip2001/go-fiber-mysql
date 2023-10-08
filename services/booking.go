package services

import (
	"time"

	"github.com/dassudip2001/webapp/models"
)

type Booking struct {
	RoomId        uint      `json:"room_id"`
	ResourceId    uint      `json:"resource_id"`
	BookStartTime time.Time `json:"book_start_time"`
	BookEndTime   time.Time `json:"book_end_time"`
	Duration      int       `json:"duration"`
}

type CreateBookingRequest struct {
	RoomId        uint      `json:"room_id"`
	ResourceId    uint      `json:"resource_id"`
	BookStartTime time.Time `json:"book_start_time"`
	BookEndTime   time.Time `json:"book_end_time"`
	Duration      int       `json:"duration"`
}

func createResponseBooking(bookingModel models.Booking) Booking {
	return Booking{
		RoomId:        bookingModel.RoomId,
		ResourceId:    bookingModel.ResourceId,
		BookStartTime: time.Time{},
		BookEndTime:   time.Time{},
		Duration:      bookingModel.Duration,
	}
}
