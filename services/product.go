package services

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
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
	var request CreateProductRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the name is not empty
	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product name is required",
		})
	}

	// Check if the price is not empty or 0

	if request.Price == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product price is required",
		})
	}

	// Check if the stock is not empty or 0

	if request.Stock == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product stock is required",
		})
	}

	// Check if the name is already in use
	var existingProduct models.Product
	result := database.Database.Db.Where("name = ?", request.Name).First(&existingProduct)
	if result.Error == nil {
		// Name already exists, return an error
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Product name   must be unique",
		})
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description, // Set Description conditionally
		Price:       request.Price,
		Stock:       request.Stock,
		// Status:      request.Status,
		CategoryID: request.CategoryID,
	}

	if err := database.Database.Db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(fiber.StatusCreated).JSON(createResponseProduct(product))
}

// get all the product
func GetAllProduct(c *fiber.Ctx) error {

	var product []models.Product
	if err := database.Database.Db.Preload("Category").Find(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to retrive the products"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// get product by id
func ProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	// check if the product is exist or not
	if database.Database.Db.Preload("Category").Find(&product, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to retrive the Poduct"})
	}

	// return  the category

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

	// check if the product is exist or not
	if database.Database.Db.First(&product, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Faild to retrive the Poduct"})
	}

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

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// check if id is valid
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product ID not provided"})
	}
	// crete a variable of type models.Product

	var product models.Product

	// check if the Product is present in the database

	if database.Database.Db.First(&product, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
	}

	// return the Product
	if err := database.Database.Db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve Product"})
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete Product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product successfully deleted"})
}
