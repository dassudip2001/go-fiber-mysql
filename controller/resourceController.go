package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

// get the resources
func GetResources(c *fiber.Ctx) error {
	return services.GetResources(c)
}

// get by id

func GetResourceById(c *fiber.Ctx) error {
	return services.GetResourcesById(c)
}

// post resources
func CreateResource(c *fiber.Ctx) error {
	return services.CreateResource(c)
}

// update resources
func UpdateResources(c *fiber.Ctx) error {
	return services.UpdateResources(c)
}

// delete resources
func DeleteResources(c *fiber.Ctx) error {
	return services.DeleteResources(c)
}
