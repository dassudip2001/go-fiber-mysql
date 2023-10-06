package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Resource struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}

type CreateResourceRequest struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}

func createResponseResource(resourceModel models.Resource) Resource {
	return Resource{
		Name:        resourceModel.Name,
		IsAvailable: resourceModel.IsAvailable,
	}
}

// create a new resource
func CreateResource(c *fiber.Ctx) error {
	var request CreateResourceRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Resource name is required",
		})
	}

	var existingResource models.Resource

	if result := database.Database.Db.Where("name=?", request.Name).First(&existingResource); result.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Resource already exists",
		})
	}

	resource := models.Resource{
		Name:        request.Name,
		IsAvailable: request.IsAvailable,
	}

	if err := database.Database.Db.Create(&resource).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Resource creation failed",
		})
	}

	responseResource := createResponseResource(resource)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Resource created successfully",
		"data":    responseResource,
	})
}

// get all resources
func GetResources(c *fiber.Ctx) error {
	var resources []models.Resource
	if err := database.Database.Db.Find(&resources).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve resources",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Resources retrieved successfully",
		"data":    resources,
	})
}

// get a resource by id
func GetResourcesById(c *fiber.Ctx) error {
	id := c.Params("id")
	var resource models.Resource

	if err := database.Database.Db.First(&resource, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Resource not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve resource",
		})
	}
	return c.JSON(resource)
}

// update resources
func UpdateResources(c *fiber.Ctx) error {
	id := c.Params("id")
	var resource models.Resource

	if err := database.Database.Db.First(&resource, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Resource not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve resource",
			"success": false,
		})
	}

	if err := c.BodyParser(&resource); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"success": false,
		})
	}

	if err := database.Database.Db.Save(&resource).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update resource",
			"success": false,
		})
	}
	return c.JSON(resource)
}

// delete resource
func DeleteResources(c *fiber.Ctx) error {
	id := c.Params("id")

	var resource models.Resource

	if err := database.Database.Db.First(&resource, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Resource not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve resource",
			"success": false,
		})
	}

	if err := database.Database.Db.Delete(&resource).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete resource",
			"success": false,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
