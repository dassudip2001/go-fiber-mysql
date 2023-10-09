package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

// create a new booking
func CreateBooking(c *fiber.Ctx) error {
	return services.CreateBooking(c)
}

// get all bookings
func GetBookings(c *fiber.Ctx) error {
	return services.GetBookings(c)
}

// get booking by id
func GetBookingById(c *fiber.Ctx) error {
	return services.GetBookingById(c)
}

// update booking
func UpdateBooking(c *fiber.Ctx) error {
	return services.UpdateBooking(c)
}

// delete booking
func DeleteBooking(c *fiber.Ctx) error {
	return services.DeleteBooking(c)
}
