package controllers

import (
	"adminProject/src/database"
	"adminProject/src/models"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {

	var users []models.User
	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}
