package routes

import (
	"api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api/v1/")
	})

	route := app.Group("/api/v1")

	route.Get("/", Root)
	route.Post("/bookmark/new", controllers.AddBookmark)
	route.Get("/bookmark/:id", controllers.GetBookmarkById)
}

// Root func
//
//	@Summary	Root URL - for health check
//	@Success	200
//	@Tags		/
//	@BasePath	/api/v1
//	@Router		/api/v1/ [get]
func Root(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
