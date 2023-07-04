package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/", root)
}

// root func
//
//	@Summary	Root URL - for health check
//	@Success	200
//	@Tags		/api/v1
//	@BasePath	/api/v1
//	@Router		/api/v1/ [get]
func root(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
