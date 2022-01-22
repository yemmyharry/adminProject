package main

import (
	"adminProject/src/database"
	"adminProject/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	routes.SetUp(app)

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}
