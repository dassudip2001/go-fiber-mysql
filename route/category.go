package route

import (
	"github.com/dassudip2001/webapp/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/v1/category", controller.CreateCategory)
	app.Get("/api/v1/category", controller.GetCategory)
	app.Get("/api/v1/category/:id", controller.GetCategoryById)
	app.Put("/api/v1/category/:id", controller.UpdateCategory)
	app.Delete("/api/v1/category/:id", controller.DeleteCategory)

}
