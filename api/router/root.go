// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package router

import "github.com/gofiber/fiber/v2"

func RunServer() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	app.Listen(":3000")
}
