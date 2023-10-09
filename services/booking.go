package services

import (
	"time"

	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		BookStartTime: bookingModel.BookStartTime,
		BookEndTime:   bookingModel.BookEndTime,
		Duration:      bookingModel.Duration,
	}
}

// create a new booking
func CreateBooking(c *fiber.Ctx) error {
	var request CreateBookingRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.RoomId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Room id is required",
		})
	}

	var existingRoom models.Booking
	result := database.Database.Db.Where("RoomId = ?", request.RoomId).First(&existingRoom)
	if result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Room Id is not valid",
		})
	}

	booking := models.Booking{
		RoomId:        request.RoomId,
		ResourceId:    request.ResourceId,
		BookStartTime: request.BookStartTime,
		BookEndTime:   request.BookEndTime,
		Duration:      request.Duration,
	}

	if err := database.Database.Db.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Booking creation failed",
		})
	}

	responseBooking := createResponseBooking(booking)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"massage": "Booking created successfully",
		"data":    responseBooking,
	})

}

// get all bookings
func GetBookings(c *fiber.Ctx) error {

	var bookings []models.Booking
	if err := database.Database.Db.Find(&bookings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"massage": "Failed to retrieve bookings",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"massage": "Bookings retrieved successfully",
		"data":    bookings,
	})
}

// get a booking by id
func GetBookingById(c *fiber.Ctx) error {

	id := c.Params("id")
	var booking models.Booking

	if err := database.Database.Db.First(&booking, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"massage": "Booking does not exists",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"massage": "Booking retrieved successfully",
		"data":    booking,
	})
}

// update a booking
func UpdateBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	var booking models.Booking

	if err := database.Database.Db.First(&booking, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"massage": "Booking not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"massage": "Failed to retrieve booking",
		})
	}
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"success": false,
		})
	}

	if err := database.Database.Db.Save(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update resource",
			"success": false,
		})
	}
	return c.JSON(booking)

}

// delete a booking
func DeleteBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	var booking models.Booking

	if err := database.Database.Db.First(&booking, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"massage": "Booking not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"massage": "Failed to retrieve booking",
		})
	}

	if err := database.Database.Db.Delete(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete booking",
			"success": false,
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"massage": "Booking deleted successfully",
	})
}
