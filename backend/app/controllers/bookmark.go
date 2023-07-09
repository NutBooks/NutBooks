package controllers

import (
	conn "api/db"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

// AddBookmark
//
//	@Summary		북마크를 DB에 추가하는 API
//	@Description	새 북마크를 DB에 저장. 북마크 링크는 필수 데이터이고, 그 외는 옵셔널.
//	@Tags			bookmark
//	@Accept			json
//	@Produce		json
//	@Param			userId		body	uint	false	"User ID"
//	@Param			title		body	string	false	"Title"
//	@Param			link		body	string	true	"Link(URL)"
//	@Param			keywords	body	string	false	"keywords"
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
