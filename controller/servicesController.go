package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

func CreateCatering(c *fiber.Ctx) error {
	return services.CreateServices(c)
}

func GetAllServices(c *fiber.Ctx) error {
	return services.GetAllServices(c)
}

func GetServiceById(c *fiber.Ctx) error {
	return services.GetServiceById(c)
}

func UpdateServices(c *fiber.Ctx) error {
	return services.UpdateServices(c)
}

func DeleteServices(c *fiber.Ctx) error {
	return services.DeleteServices(c)
}
