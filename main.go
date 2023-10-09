package main

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")
	// initialized the database connection
	database.ConnectDb()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//api routes setup
	route.SetupRoutes(app)
	// web routes setup
	route.SetupWebRoutes(app)
	// product routes setup
	route.ProductRouter(app)
	// location routes setup
	route.LocationRouter(app)

	// add the metting room routes
	route.RouteMettingRoom(app)
	// add resources router
	route.ResourceRouter(app)
	// add booking router
	route.BookingRoute(app)

	// load the static files
	app.Static("/", "./public")

	app.Listen(":3000")
}
