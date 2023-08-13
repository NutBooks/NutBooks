package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// AddBookmarkHandler
//
//	@Summary		특정 유저가 북마크를 DB에 추가하는 API
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
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
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
	log.Infow("[func AddBookmarkHandler]", "params", params)

	bookmark := &models.Bookmark{
		UserID: params.UserID,
		Title:  params.Title,
		Link:   params.Link,
	}
	bookmark, err = crud.AddBookmark(bookmark)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}
	log.Debugw("[func AddBookmarkHandler]", "bookmark", bookmark)

	return c.Status(fiber.StatusOK).JSON(models.AddBookmarkResponse{
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
//	@Success	200	{object}	models.GetBookmarkByIdResponse{data=models.Bookmark}
//	@Failure	400	{object}	models.GetBookmarkByIdResponse{}
//	@Router		/api/v1/bookmark/{id}/ [get]
func GetBookmarkByIdHandler(c *fiber.Ctx) error {
	params := models.Bookmark{}
	err := c.ParamsParser(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validator := &utils.Validator{}
	validateError := validator.Validate(params)
	if validateError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Data:    validateError,
			Message: "Validation failed",
		})
	}
	log.Infow("[func GetBookmarkByIdHandler]", "params", params)

	found, err := crud.GetBookmarkById(params.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func GetAllBookmarksHandler]", "params", params)

	found, err := crud.GetAllBookmarks(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}
	log.Debugw("[func GetAllBookmarksHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetAllBookmarksResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
