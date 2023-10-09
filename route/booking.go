package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func BookingRoute(c *fiber.App) {
	c.Get("/api/v1/booking", controller.GetBookings)
	c.Get("/api/v1/booking/:id", controller.GetBookingById)
	c.Post("/api/v1/booking", controller.CreateBooking)
	c.Put("/api/v1/booking/:id", controller.UpdateBooking)
	c.Delete("/api/v1/booking/:id", controller.DeleteBooking)
}
