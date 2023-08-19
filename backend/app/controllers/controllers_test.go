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
	user.Post("", AddUserHandler)
	user.Get("/:id<int>", GetUserByIdHandler)
	user.Get("", GetAllUsersHandler)
	user.Get("/check-email", CheckEmailDuplicateHandler)

	// Auth router
	auth := route.Group("/auth")
	auth.Post("/login", LogInHandler)

	// bookmark
	bookmark := route.Group("/bookmark")
	bookmark.Post("/", AddBookmarkHandler)
	bookmark.Get("/:id/", GetBookmarkByIdHandler)
	bookmark.Get("/", GetAllBookmarksHandler)
}

func TestControllers(t *testing.T) {
	t.Helper()
	controllersTestApp(t)

	t.Run("testUserController", testUserController)
	t.Run("testAuthenticationController", testAuthenticationController)
	t.Run("testBookmarkController", testBookmarkController)
}
