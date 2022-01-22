package routes

import (
	controllers2 "adminProject/src/controllers"
	"adminProject/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Post("/admin/register", controllers2.Register)
		api.Post("/admin/login", controllers2.Login)

		authenticatedRoute := api.Use(middlewares.IsAuthenticated)
		{
			authenticatedRoute.Get("/admin/user", controllers2.User)
			authenticatedRoute.Post("/admin/logout", controllers2.Logout)
			authenticatedRoute.Put("/admin/update_user", controllers2.UpdateUser)
			authenticatedRoute.Get("/admin/ambassador", controllers2.Ambassadors)
			authenticatedRoute.Get("/admin/products", controllers2.Products)
			authenticatedRoute.Post("/admin/products", controllers2.CreateProduct)
			authenticatedRoute.Get("/admin/products/:id", controllers2.GetProduct)
			authenticatedRoute.Put("/admin/products/:id", controllers2.UpdateProduct)
			authenticatedRoute.Delete("/admin/products/:id", controllers2.DeleteProduct)
			authenticatedRoute.Get("/admin/users/:id/links", controllers2.Link)
			authenticatedRoute.Get("/admin/orders", controllers2.Orders)
		}
	}
}
