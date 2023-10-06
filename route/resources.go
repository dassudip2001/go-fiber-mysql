package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func ResourceRouter(c *fiber.App) {
	c.Get("/api/v1/resources", controller.GetResources)
	c.Post("/api/v1/resources", controller.CreateResource)
	c.Get("/api/v1/resources/:id", controller.GetResourceById)
	c.Put("/api/v1/resources/:id", controller.UpdateResources)
	c.Delete("/api/v1/resources/:id", controller.DeleteResources)

}
