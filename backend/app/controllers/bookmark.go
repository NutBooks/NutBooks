package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
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

	err := c.BodyParser(params)
	if err != nil {
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

	bookmark := &models.Bookmark{
		UserID: params.UserID,
		Title:  params.Title,
		Link:   params.Link,
	}

	err = crud.AddBookmark(bookmark)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddBookmarkResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.AddBookmarkResponse{
		Error:   false,
		Message: "Success",
		Data:    bookmark,
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

	found, err := crud.GetBookmarkById(param.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetBookmarkByIdResponse{
			Error:   true,
			Message: err.Error(),
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
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	found, err := crud.GetAllBookmarks(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllBookmarksResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetAllBookmarksResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
