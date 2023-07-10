package controllers

import (
	conn "api/db"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

type GetBookmarkByIdJSONResult struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetBookmarkById
//
//	@Summary	ID를 사용해 북마크 1개 정보 읽기
//	@Tags		bookmark
//	@Produce	json
//	@Param		bookmark_id	path		uint	true	"Bookmark ID"
//	@Success	200			{object}	GetBookmarkByIdJSONResult
//	@Failure	400
//	@BasePath	/api/v1
//	@Router		/api/v1/bookmark/{id} [get]
func GetBookmarkById(c *fiber.Ctx) error {

	param := models.Bookmark{}

	err := c.ParamsParser(&param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	log.Println("bookmark_id: ", param.ID)

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var found models.Bookmark
	result := db.First(&found, param.ID)
	if result == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": result.Error.Error(),
		})
	}

	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   true,
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Success",
		"data":    found,
	})
}
