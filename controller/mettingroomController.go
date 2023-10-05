package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

// create a new meting room controller
func CreateMettingRoom(c *fiber.Ctx) error {
	return services.CreateMettingRoom(c)
}

// get all metting rooms controller
func GetMettingRoom(c *fiber.Ctx) error {
	return services.GetMettingRooms(c)
}

// get a single metting room controller
func GetMettingRoomById(c *fiber.Ctx) error {
	return services.GetMettingRoomById(c)
}

// update a metting room controller
func UpdateMettingRoom(c *fiber.Ctx) error {
	return services.UpdateMettingRoom(c)
}

// delete a metting room controller
func DeleteMettingRoom(c *fiber.Ctx) error {
	return services.DeleteMettingRoom(c)
}
