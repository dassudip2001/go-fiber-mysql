package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// The `createResponsecategory` function is a helper function that takes a `models.category` object as input
// and creates a new `category` object with the same values. It is used to convert a `models.category` object
// to a `category` object, which is a simplified version of the category model that will be returned as a JSON
// response.
func createResponsecategory(categoryModel models.Category) Category {
	return Category{
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
	}
}

// create a category
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Do not specify the ID field here
	if err := database.Database.Db.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
	}

	responsecategory := createResponsecategory(category)
	return c.Status(fiber.StatusCreated).JSON(responsecategory)
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
	id := c.Params("id")

	var category models.Category
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
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category ID not provided"})
	}

	var category models.Category
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
