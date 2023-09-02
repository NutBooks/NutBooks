package middlewares

import (
	"api/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware(app *fiber.App) {
	origins := configs.AllowOrigins
	if configs.FeDevHost != "" {
		origins += ", " + configs.FeDevHost
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		MaxAge:       600,
	}))
}
