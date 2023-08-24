package routes

import (
	_ "api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {
	// Swagger docs
	// https://github.com/gofiber/swagger
	// https://github.com/swaggo/swag

	route := app.Group("/docs")

	route.Get("*", swagger.HandlerDefault)
}
