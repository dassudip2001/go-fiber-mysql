package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

// create a new category controller
func CreateLocation(c *fiber.Ctx) error {
	return services.CreateLocation(c)
}

// get all categories controller
func GetLocation(c *fiber.Ctx) error {
	return services.GetLocations(c)
}

// get a single category controller
func GetLocationById(c *fiber.Ctx) error {
	return services.GetLocationById(c)
}

// update a category controller
func UpdateLocation(c *fiber.Ctx) error {
	return services.UpdateLocation(c)
}

// delete a category controller
func DeleteLocation(c *fiber.Ctx) error {
	return services.DeleteLocation(c)
}
