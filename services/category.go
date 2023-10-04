package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Category struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// Define a function to create a response category
func createResponseCategory(categoryModel models.Category) Category {
	return Category{
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
	}
}

func CreateCategory(c *fiber.Ctx) error {
	var request CreateCategoryRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
		})
	}

	// Check if the name is already in use
	var existingCategory models.Category
	result := database.Database.Db.Where("name = ?", request.Name).First(&existingCategory)
	if result.Error == nil {
		// Name already exists, return an error
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Category name must be unique",
		})
	}

	category := models.Category{
		Name:        request.Name,
		Description: request.Description, // Set Description conditionally
	}

	if err := database.Database.Db.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
	}

	responseCategory := createResponseCategory(category)
	return c.Status(fiber.StatusCreated).JSON(responseCategory)
}

// get all the category
func GetCategory(c *fiber.Ctx) error {
	var category []models.Category

	if err := database.Database.Db.Find(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories"})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// get category by id
func GetCategoryById(c *fiber.Ctx) error {
	// get the id from the url
	id := c.Params("id")

	// crete a variable of type models.Category
	var category models.Category
	// check if the category is present in the database

	if database.Database.Db.First(&category, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
	}
	// return the category
	if err := database.Database.Db.Find(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// update the category
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	if database.Database.Db.First(&category, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
	}

	if err := database.Database.Db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := database.Database.Db.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// delete the category

func DeleteCategory(c *fiber.Ctx) error {
	// get the id from the url
	id := c.Params("id")

	// check if id is valid
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category ID not provided"})
	}
	// crete a variable of type models.Category

	var category models.Category

	// check if the category is present in the database

	if database.Database.Db.First(&category, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
	}

	// return the category
	if err := database.Database.Db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
	}

	if err := database.Database.Db.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete category"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category successfully deleted"})
}
