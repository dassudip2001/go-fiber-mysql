package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryID  uint   `json:"category_id"`
}

func createResponseProduct(productModel models.Product) Product {
	return Product{
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
		CategoryID:  productModel.CategoryID,
	}
}

// create a new product

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body"})
	}

	if err := database.Database.Db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create Product"})
	}
	responseproduct := createResponseProduct(product)
	return c.Status(fiber.StatusCreated).JSON(responseproduct)
}

// get all the product
func GetAllProduct(c *fiber.Ctx) error {

	var product []models.Product
	if err := database.Database.Db.Find(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to retrive the products"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// get product by id
func ProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if err := database.Database.Db.Find(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to retrive the Poduct"})

	}
	return c.Status(fiber.StatusOK).JSON(product)
}

// update the product
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := database.Database.Db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not Found Asscieted this id"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to Update the product"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := database.Database.Db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// delete the product
