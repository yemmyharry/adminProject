package controllers

import (
	"adminProject/src/database"
	"adminProject/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.TotalOrder()
	}

	return c.JSON(orders)

}
