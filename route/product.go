package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(c *fiber.App) {
	c.Get("/api/v1/products", controller.GetAllProducts)
	c.Post("/api/v1/products", controller.CreateProduct)

}
