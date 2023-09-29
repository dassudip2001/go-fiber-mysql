package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	// Request handling code
	return services.CreateCategory(c)
}

func GetCategory(c *fiber.Ctx) error {
	// Request handling code
	return services.GetCategory(c)
}

func GetCategoryById(c *fiber.Ctx) error {
	// Request handling code
	return services.GetCategoryById(c)
}

func UpdateCategory(c *fiber.Ctx) error {
	// Request handling code
	return services.UpdateCategory(c)
}

func DeleteCategory(c *fiber.Ctx) error {
	// Request handling code
	return services.DeleteCategory(c)
}
