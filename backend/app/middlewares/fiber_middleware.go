package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		// https://docs.gofiber.io/api/middleware/logger/
		logger.New(logger.Config{
			TimeFormat: "2023-06-27T11:00:00Z",
			TimeZone:   "UTC",
		}),
	)
}
