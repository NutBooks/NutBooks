package controllers

import (
	conn "api/db"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

type AddBookmarkJsonRequest struct {
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Link   string `json:"link" example:"https://cheesecat47.github.io"`
}

// AddBookmark
//
//	@Summary		북마크를 DB에 추가하는 API
//	@Description	새 북마크를 DB에 저장. 북마크 링크는 필수 데이터이고, 그 외는 옵셔널.
//	@Tags			bookmark
//	@Accept			json
//	@Produce		json
//	@Param			request	body	AddBookmarkJsonRequest	true	"body params"
//	@Success		200
//	@Failure		400
//	@BasePath		/api/v1
//	@Router			/api/v1/bookmark/new [post]
func AddBookmark(c *fiber.Ctx) error {
	// Get claims from JWT
	// Check user permissions to create a new bookmark

	bookmark := &models.Bookmark{}

	if err := c.BodyParser(bookmark); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	log.Println("bookmark: ", bookmark)

	if bookmark.Link == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "link is required parameter",
		})
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	result := db.Create(&bookmark)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": result.Error.Error(),
		})
	}
	log.Println(result)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Success",
	})
}
