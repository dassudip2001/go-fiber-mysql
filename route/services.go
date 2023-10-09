package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func ServicesRoute(c *fiber.App) {
	c.Get("/api/v1/services", controller.GetAllServices)
	c.Get("/api/v1/services/:id", controller.GetServiceById)
	c.Post("/api/v1/services", controller.CreateCatering)
	c.Put("/api/v1/services/:id", controller.UpdateServices)
	c.Delete("/api/v1/services/:id", controller.DeleteServices)
}
