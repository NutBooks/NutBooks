// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "api/docs"
)

func RunServer() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	// Swagger docs
	// https://github.com/gofiber/swagger
	// https://github.com/swaggo/swag
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
