package routes

import (
	"api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api/v1")
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", Root)

	// Bookmark router
	bookmark := v1.Group("/bookmark")
	bookmark.Post("/", controllers.AddBookmark)
	bookmark.Get("/:id", controllers.GetBookmarkById)
	bookmark.Get("/", controllers.GetAllBookmarks)

	// User router
	user := v1.Group("/user")
	user.Post("/", controllers.AddUser)
	user.Get("/:id", controllers.GetUserById)
	user.Get("/", controllers.GetAllUsers)
}

// Root func
//
//	@Summary	Root URL - for health check
//	@Success	200
//	@Tags		/
//	@Router		/api/v1/ [get]
func Root(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
