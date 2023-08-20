package routes

import (
	"api/app/controllers"
	"api/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api/v1")
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", Root)

	// Bookmark router
	{
		// 북마크 관련 API는 모두 유저 로그인 상태 (유저 아이디) 필요
		bookmark := v1.Group("/bookmark")
		bookmark.Post("", middlewares.Protected(), controllers.AddBookmarkHandler)
		bookmark.Get("/:id<int>", middlewares.Protected(), controllers.GetBookmarkByIdHandler)
		bookmark.Get("", middlewares.Protected(), controllers.GetAllBookmarksHandler)
	}

	// User router
	{
		user := v1.Group("/user")
		user.Post("", controllers.AddUserHandler)
		user.Get("/check-email", controllers.CheckEmailDuplicateHandler)

		// 아래 두 엔드포인트는 추후 관리자 인가도 검증 필요
		user.Get("/:id<int>", middlewares.Protected(), controllers.GetUserByIdHandler)
		user.Get("", middlewares.Protected(), controllers.GetAllUsersHandler)
	}

	// Auth router
	{
		auth := v1.Group("/auth")
		auth.Post("/login", controllers.LogInHandler)
	}
}

// Root func
//
//	@Summary	Root URL - for health check
//	@Success	200
//	@Tags		/
//	@Router		/api/v1/ [get]
func Root(c *fiber.Ctx) error {
	return c.SendString("Hi")
}
