package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func LocationRouter(c *fiber.App) {

	c.Get("/api/v1/locations", controller.GetLocation)
	c.Post("/api/v1/locations", controller.CreateLocation)
	c.Get("/api/v1/locations/:id", controller.GetLocationById)
	c.Put("/api/v1/locations/:id", controller.UpdateLocation)
	c.Delete("/api/v1/locations/:id", controller.DeleteLocation)

}
