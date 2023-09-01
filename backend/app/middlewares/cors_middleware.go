package middlewares

import (
	"api/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware(app *fiber.App) {
	origins := "https://nutbooks.koreacentral.cloudapp.azure.com"
	if configs.FeDevHost != "" {
		origins += ", " + configs.FeDevHost
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		MaxAge:       600,
	}))

	app.Use(func(c *fiber.Ctx) error {
		h := c.GetReqHeaders()
		if configs.FeDevHost != "" && configs.CorsSecret != "" {
			if h["Origin"] == configs.FeDevHost && h["X-Custom-Origin-Token"] == configs.CorsSecret {
				return c.Next()
			}
		}
		c.Response().Header.Set("Access-Control-Allow-Origin", "")
		return c.SendStatus(fiber.StatusNoContent)
	})
}
