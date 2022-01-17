package routes

import (
	"adminProject/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Post("/admin/register", controllers.Register)

	}
}