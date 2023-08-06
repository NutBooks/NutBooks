package controllers

import (
	"api/app/utils"
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AddBookmarkHandler
//
//	@Summary		북마크를 DB에 추가하는 API
//	@Description	새 북마크를 DB에 저장. 북마크 링크는 필수 데이터이고, 그 외는 옵셔널.
//	@Tags			bookmark
//	@Accept			json
//	@Produce		json
//	@Param			params	body		models.AddBookmarkRequest{}	true	"body params"
//	@Success		200		{object}	models.AddBookmarkResponse{data=models.Bookmark}
//	@Failure		400		{object}	models.AddBookmarkResponse{}
//	@Router			/api/v1/bookmark/ [post]
func AddBookmarkHandler(c *fiber.Ctx) error {
	// Get claims from JWT
	// Check user permissions to create a new bookmark

	params := &models.AddBookmarkRequest{}

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	bookmark := &models.Bookmark{
		UserID: params.UserID,
		Title:  params.Title,
		Link:   params.Link,
	}

	result := db.Create(bookmark)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.AddBookmarkResponse{
		Error:   false,
		Message: "Success",
		Data:    nil,
	})
}

// GetBookmarkByIdHandler
//
//	@Summary	ID를 사용해 북마크 1개 정보 읽기
//	@Tags		bookmark
//	@Produce	json
//	@Param		id	path		uint	true	"Bookmark ID"
//	@Success	200	{object}	models.GetBookmarkByIdResponse{data=models.Bookmark}
//	@Failure	400	{object}	models.GetBookmarkByIdResponse{}
//	@Router		/api/v1/bookmark/{id}/ [get]
func GetBookmarkByIdHandler(c *fiber.Ctx) error {

	param := models.Bookmark{}

	err := c.ParamsParser(&param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var found models.Bookmark
	result := db.First(&found, param.ID)
	if result == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: "Cannot find this bookmark",
			Data:    nil,
		})
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetBookmarkByIdResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}

// GetAllBookmarksHandler
//
//	@Summary	offset부터 limit까지 북마크 목록을 반환
//	@Tags		bookmark
//	@Produce	json
//	@Param		offset	query		int	false	"limit과 offset은 같이 입력해야 합니다"
//	@Param		limit	query		int	false	"limit과 offset은 같이 입력해야 합니다"
//	@Success	200		{object}	models.GetAllBookmarksResponse{data=[]models.Bookmark{}}
//	@Failure	400		{object}	models.GetAllBookmarksResponse{}
//	@Router		/api/v1/bookmark/ [get]
func GetAllBookmarksHandler(c *fiber.Ctx) error {
	params := &models.GetAllBookmarksRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var found []models.Bookmark
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = db.Find(&found)
	} else {
		result = db.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}
	if result == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: "Cannot find any bookmarks",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetAllBookmarksResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
