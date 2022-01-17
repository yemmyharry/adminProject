package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	return c.JSON(fiber.Map{
		"message": "Registered successfully",
	})
}
