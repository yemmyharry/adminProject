package routes

import (
	"adminProject/controllers"
	"adminProject/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Post("/admin/register", controllers.Register)
		api.Post("/admin/login", controllers.Login)

		authenticatedRoute := api.Use(middlewares.IsAuthenticated)
		{
			authenticatedRoute.Get("/admin/user", controllers.User)
			authenticatedRoute.Post("/admin/logout", controllers.Logout)
			authenticatedRoute.Put("/admin/update_user", controllers.UpdateUser)
			authenticatedRoute.Get("/admin/ambassador", controllers.Ambassadors)
			authenticatedRoute.Get("/admin/products", controllers.Products)
			authenticatedRoute.Post("/admin/products", controllers.CreateProduct)
		}
	}
}
