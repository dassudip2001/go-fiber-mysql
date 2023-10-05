package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func RouteMettingRoom(c *fiber.App) {

	c.Get("/api/v1/mettingrooms", controller.GetMettingRoom)
	c.Post("/api/v1/mettingrooms", controller.CreateMettingRoom)
	c.Get("/api/v1/mettingrooms/:id", controller.GetMettingRoomById)
	c.Put("/api/v1/mettingrooms/:id", controller.UpdateMettingRoom)
	c.Delete("/api/v1/mettingrooms/:id", controller.DeleteMettingRoom)

}
