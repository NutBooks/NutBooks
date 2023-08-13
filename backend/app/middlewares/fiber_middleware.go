package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		// https://docs.gofiber.io/api/middleware/logger/
		logger.New(logger.Config{
			TimeFormat: time.RFC3339,
			TimeZone:   "UTC",
		}),
	)
}
