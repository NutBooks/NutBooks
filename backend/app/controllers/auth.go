package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
)

// LogInHandler
//
//	@Summary		로그인 API
//	@Description	비밀번호는 영문 + 숫자 8-12자리
//	@Tags			auth
//	@Produce		json
//	@Param			params	body		models.LogInRequest	true	"body params"
//	@Success		200		{object}	models.LogInResponse{}
//	@Failure		400		{object}	models.LogInResponse{}
//	@Failure		401		{object}	models.LogInResponse{}
//	@Failure		500		{object}	models.LogInResponse{}
//	@Router			/api/v1/auth/login/ [post]
func LogInHandler(c *fiber.Ctx) error {
	params := &models.LogInRequest{}
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	user, err := crud.GetUserByEmail(params.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	password, err := crud.GetPasswordByUserId(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	if params.Password != password.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(models.LogInResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.LogInResponse{
		Error:   false,
		Message: "Success",
		Data:    nil,
	})
}
