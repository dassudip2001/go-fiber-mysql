package controller

import (
	"github.com/dassudip2001/webapp/services"
	"github.com/gofiber/fiber/v2"
)

// get all the product

func GetAllProducts(c *fiber.Ctx) error {
	return services.GetAllProduct(c)
}

// create the procuct

func CreateProduct(c *fiber.Ctx) error {
	return services.CreateProduct(c)
}

// get by id

func GetProductById(c *fiber.Ctx) error {
	return services.ProductById(c)
}

// update

func UpdateProduct(c *fiber.Ctx) error {
	return services.UpdateProduct(c)
}

// delete
func DeleteProduct(c *fiber.Ctx) error {
	return services.DeleteProduct(c)
}
