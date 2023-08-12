package controllers

import (
	"api/app/middlewares"
	"api/configs"
	conn "api/db"
	"github.com/gofiber/fiber/v2"
	"testing"
)

var (
	app *fiber.App
)

func controllersTestApp(t *testing.T) {
	config := configs.FiberConfig()
	conn.Connect()
	app = fiber.New(config)
	middlewares.FiberMiddleware(app)

	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Post("/", AddUserHandler)
	user.Get("/:id/", GetUserByIdHandler)
	user.Get("/", GetAllUsersHandler)

	// Auth router
	auth := route.Group("/auth")
	auth.Post("/login/", LogInHandler)
}

func TestControllers(t *testing.T) {
	t.Helper()
	controllersTestApp(t)

	t.Run("testUserController", testUserController)
	t.Run("testAuthenticationController", testAuthenticationController)
}
