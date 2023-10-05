package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateMettingRoomRequest struct {
	Name          string `json:"name"`
	Capacity      int    `json:"capacity"`
	Configuration string `json:"configuration"`
	IsAvailable   bool   `json:"is_available"`
	LocationId    uint   `json:"location_id"`
}

type MettingRoom struct {
	Name          string `json:"name"`
	Capacity      int    `json:"capacity"`
	Configuration string `json:"configuration"`
	IsAvailable   bool   `json:"is_available"`
	LocationId    uint   `json:"location_id"`
}

func createResponseMettingRoom(mettingRoomModel models.MettingRoom) MettingRoom {
	return MettingRoom{
		Name:          mettingRoomModel.Name,
		Capacity:      mettingRoomModel.Capacity,
		Configuration: mettingRoomModel.Configuration,
		IsAvailable:   mettingRoomModel.IsAvailable,
		LocationId:    mettingRoomModel.LocationId,
	}
}

// create a new mettingRoom
func CreateMettingRoom(c *fiber.Ctx) error {
	var request CreateMettingRoomRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "MettingRoom name is required",
		})
	}

	if request.Capacity == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "MettingRoom capacity is required",
		})
	}

	if request.Configuration == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "MettingRoom configuration is required",
		})
	}

	if request.LocationId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "MettingRoom location is required",
		})
	}

	// Check if the name is already in use

	var existingMettingRoom models.MettingRoom

	result := database.Database.Db.Where("name = ?", request.Name).First(&existingMettingRoom)
	if result.Error == nil {
		// Name already exists, return an error
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "MettingRoom name must be unique",
		})
	}

	mettingRoom := models.MettingRoom{
		Name:          request.Name,
		Capacity:      request.Capacity,
		Configuration: request.Configuration,
		IsAvailable:   request.IsAvailable,
		LocationId:    request.LocationId,
	}

	// Create the mettingRoom
	if err := database.Database.Db.Create(&mettingRoom).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not create mettingRoom",
		})
	}
	responceMettingRoom := createResponseMettingRoom(mettingRoom)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "MettingRoom created successfully",
		"data":    responceMettingRoom,
	})
}

// Get all mettingRooms
func GetMettingRooms(c *fiber.Ctx) error {
	var mettingRooms []models.MettingRoom

	database.Database.Db.Preload("Location").Find(&mettingRooms)
	return c.JSON(fiber.Map{
		"success": true,
		"data":    mettingRooms,
	})

}

// Get a single mettingRoom
func GetMettingRoomById(c *fiber.Ctx) error {
	id := c.Params("id")
	var mettingRoom models.MettingRoom

	if database.Database.Db.Preload("Location").First(&mettingRoom, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not find mettingRoom",
		})
	}

	if err := database.Database.Db.Preload("location").Find(&mettingRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "MettingRoom not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not retrieve mettingRoom",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    mettingRoom,
	})
}

// Update a mettingRoom
func UpdateMettingRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var mettingRoom models.MettingRoom
	if database.Database.Db.First(&mettingRoom, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not find mettingRoom",
		})
	}

	if err := database.Database.Db.First(&mettingRoom, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "MettingRoom not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not retrieve mettingRoom",
		})
	}

	if err := c.BodyParser(&mettingRoom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid body request",
			"success": false,
		})
	}

	if err := database.Database.Db.Save(&mettingRoom).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to update location",
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "MettingRoom updated successfully",
		"data":    mettingRoom,
	})
}

// delete mettingRoom
func DeleteMettingRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	var mettingRoom models.MettingRoom

	if database.Database.Db.First(&mettingRoom, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrive location",
			"success": false,
		})
	}

	if err := database.Database.Db.First(&mettingRoom, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "location not found",
				"Success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrive Location",
			"success": false,
		})
	}

	if err := database.Database.Db.Delete(&mettingRoom).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete location",
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Location deleted successfully",
	})
}
