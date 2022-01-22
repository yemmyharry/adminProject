package controllers

import (
	"adminProject/src/database"
	"adminProject/src/models"
	"github.com/gofiber/fiber/v2"
)

func Link(c *fiber.Ctx) error {
	id := c.Params("id")
	var links []models.Link
	database.DB.Where("id = ?", id).Find(&links)
	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)
		links[i].Orders = orders
	}
	return c.JSON(links)

}
