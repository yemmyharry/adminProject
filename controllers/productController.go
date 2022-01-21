package controllers

import (
	"adminProject/src/database"
	"adminProject/src/models"
	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(503).JSON(err)
	}
	database.DB.Create(&product)
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.First(&product, id)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.First(&product, id)
	if err := c.BodyParser(&product); err != nil {
		return c.Status(503).JSON(err)
	}
	database.DB.Save(&product)
	return c.JSON(product)
}

// DeleteProduct deletes a product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.First(&product, id)
	database.DB.Delete(&product)
	return c.JSON(fiber.Map{"status": "successfully deleted"})
}
