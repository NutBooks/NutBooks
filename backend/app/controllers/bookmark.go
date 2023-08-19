package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddBookmarkHandler
//
//	@Summary		특정 유저가 북마크를 DB에 추가하는 API
//	@Description	새 북마크를 DB에 저장. 북마크 링크는 필수 데이터이고, 그 외는 옵셔널.
//	@Tags			bookmark
//	@Accept			json
//	@Produce		json
//	@Param			params	body		models.AddBookmarkRequest	true	"body params"
//	@Success		201		{object}	models.AddBookmarkResponse
//	@Failure		400		{object}	models.AddBookmarkWithErrorResponse
//	@Failure		500		{object}	models.AddBookmarkWithErrorResponse
//	@Router			/api/v1/bookmark [post]
func AddBookmarkHandler(c *fiber.Ctx) error {
	// Get claims from JWT
	// Check user permissions to create a new bookmark

	params := &models.AddBookmarkRequest{}
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkWithErrorResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func AddBookmarkHandler]", "params", params)

	bookmark := &models.Bookmark{
		UserID: params.UserID,
		Title:  params.Title,
		Link:   params.Link,
	}
	bookmark, err = crud.AddBookmark(bookmark)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddBookmarkWithErrorResponse{
			Error:   true,
			Message: "Failed to create bookmark",
			Data:    err,
		})
	}
	log.Debugw("[func AddBookmarkHandler]", "bookmark", bookmark)

	return c.Status(fiber.StatusCreated).JSON(models.AddBookmarkResponse{
		Error:   false,
		Message: "Success",
		Data:    bookmark,
	})
}

// GetBookmarkByIdHandler
//
//	@Summary	특정 유저가 저장한 북마크 중 id가 일치하는 북마크 상세 정보 1개를 반환
//	@Tags		bookmark
//	@Produce	json
//	@Param		id	path		uint	true	"Bookmark ID"
//	@Success	200	{object}	models.GetBookmarkByIdResponse
//	@Failure	400	{object}	models.GetBookmarkByIdWithErrorResponse
//	@Failure	500	{object}	models.GetBookmarkByIdWithErrorResponse
//	@Router		/api/v1/bookmark/{id} [get]
func GetBookmarkByIdHandler(c *fiber.Ctx) error {
	// check authentication

	params := &models.GetBookmarkByIdRequest{}
	err := c.ParamsParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateError := validator.Validate(params)
	if validateError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateError,
		})
	}
	log.Infow("[func GetBookmarkByIdHandler]", "params", params)

	found, err := crud.GetBookmarkById(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdWithErrorResponse{
				Error:   true,
				Message: "record not found",
				Data:    err,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GetBookmarkByIdWithErrorResponse{
				Error:   true,
				Message: "Failed to get bookmark",
				Data:    err,
			})
		}
	}
	log.Debugw("[func GetBookmarkByIdHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetBookmarkByIdResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}

// GetAllBookmarksHandler
//
//	@Summary	특정 유저가 저장한 북마크 중 offset부터 limit까지 목록을 반환
//	@Tags		bookmark
//	@Produce	json
//	@Param		offset	query		int	false	"특정 id부터 조회할 때 사용"
//	@Param		limit	query		int	false	"limit 개수만큼 조회할 때 사용"
//	@Success	200		{object}	models.GetAllBookmarksResponse
//	@Failure	400		{object}	models.GetAllBookmarksWithErrorResponse
//	@Failure	500		{object}	models.GetAllBookmarksWithErrorResponse
//	@Router		/api/v1/bookmark [get]
func GetAllBookmarksHandler(c *fiber.Ctx) error {
	// check authentication

	params := &models.GetAllBookmarksRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func GetAllBookmarksHandler]", "params", params)

	found, err := crud.GetAllBookmarks(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksWithErrorResponse{
			Error:   true,
			Message: "Failed to get bookmarks",
			Data:    err,
		})
	}
	log.Debugw("[func GetAllBookmarksHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetAllBookmarksResponse{
		Error:   false,
		Message: "Success",
		Data: struct {
			Data []models.Bookmark `json:"data"`
			Size int               `json:"size"`
		}{
			Data: found,
			Size: len(found),
		},
	})
}
