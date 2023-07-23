package controllers

import (
	"api/app/utils"
	conn "api/db"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AddUser
//
//	@Summary	새 유저를 추가하는 API
//	@Tags		user
//	@Produce	json
//	@Param		params	body		models.AddUserRequest	true	"body params"
//	@Success	200		{object}	models.AddUserResponse{data=models.User}
//	@Failure	400		{object}	models.AddUserResponse{}
//	@Failure	500		{object}	models.AddUserResponse{}
//	@Router		/api/v1/user/ [post]
func AddUser(c *fiber.Ctx) error {
	params := &models.AddUserRequest{}

	err := c.BodyParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
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

	user := &models.User{
		Name:      params.Name,
		Authority: models.AuthorityNone,
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result := db.Create(user)
	if result == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.AddUserResponse{
		Error:   false,
		Message: "Success",
		Data:    user,
	})
}

// GetUserById
//
//	@Summary	UserID를 사용해 유저 1명 정보 읽기
//	@Tags		user
//	@Produce	json
//	@Param		id	path		uint	true	"User ID"
//	@Success	200	{object}	models.GetUserByIdResponse{data=models.User}
//	@Failure	400	{object}	models.GetUserByIdResponse{}
//	@Failure	500	{object}	models.AddUserResponse{}
//	@Router		/api/v1/user/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	params := &models.GetUserByIdRequest{}

	err := c.ParamsParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	found := &models.User{}
	result := db.First(found, params.ID)
	if result == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusOK).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetUserByIdResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
