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
//	@Param			params	formData	models.LogInRequest	true	"formData"
//	@Success		200		{object}	models.LogInResponse{}
//	@Failure		400		{object}	models.LogInResponse{}
//	@Router			/api/v1/auth/login/ [post]
func LogInHandler(c *fiber.Ctx) error {
	params := &models.LogInRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
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

	found, err := crud.GetUserIdByEmail(params.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	_, err = crud.GetUserById(found.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	// auth.Password 체크 필요

	return c.Status(fiber.StatusOK).JSON(models.LogInResponse{
		Error:   false,
		Message: "Success",
		Data:    nil,
	})
}
