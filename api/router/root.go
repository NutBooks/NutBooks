// ----------------------------------------------------------------
//
// API Server
// https://docs.gofiber.io/
//
// ----------------------------------------------------------------

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"log"
	"time"

	_ "api/docs"
)

func RunServer() {

	app := fiber.New()

	// limit 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	// Swagger docs
	// https://github.com/gofiber/swagger
	// https://github.com/swaggo/swag
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
