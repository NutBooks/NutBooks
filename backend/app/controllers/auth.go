package controllers

import (
	"api/app/utils"
	"api/configs"
	"api/db/crud"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

// LogInHandler
//
//	@Summary		로그인 API
//	@Description	로그인 성공 시 200 "Success" 메시지 반환.
//	@Description	이메일 문제 시 400 "Email not found", 비밀번호 문제 시 "Failed to login" 반환.
//	@Description	로그인 중 서버 문제 발생 시 "Failed to check ***" 반환.
//	@Tags			auth
//	@Produce		json
//	@Param			params	body		models.LogInRequest	true	"비밀번호는 영문 + 숫자 8-12자리"
//	@Success		200		{object}	models.LogInResponse
//	@Failure		400		{object}	models.LogInWithErrorResponse
//	@Failure		401		{object}	models.LogInWithErrorResponse
//	@Failure		500		{object}	models.LogInWithErrorResponse
//	@Router			/auth/login [post]
func LogInHandler(c *fiber.Ctx) error {
	params := &models.LogInRequest{}
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func LogInHandler]", "params", params)

	user, err := crud.GetUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "[func GetUserByEmail] Cannot find this user" || err.Error() == "[func GetUserByEmail] Cannot find this user" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.LogInWithErrorResponse{
				Error:   true,
				Message: "Email not found",
				Data:    err,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.LogInWithErrorResponse{
				Error:   true,
				Message: "Failed to check email",
				Data:    err,
			})
		}
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "User not found",
			Data:    err,
		})
	}

	password, err := crud.GetPasswordByUserId(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "[func GetPasswordByUserId] Cannot find this password" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.LogInWithErrorResponse{
				Error:   true,
				Message: "Failed to login",
				Data:    err,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.LogInWithErrorResponse{
				Error:   true,
				Message: "Failed to check password",
				Data:    err,
			})
		}
	}
	if password == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "Failed to login",
			Data:    err,
		})
	}
	// 현재는 PW가 평문으로 오는데 추후 hashing 추가하면 여기서 확인
	if params.Password != password.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "Failed to login",
			Data:    err,
		})
	}

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"authority": user.Authority,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
		"iss":       "cheesecat47@gmail.com@Nutbooks",
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(configs.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.LogInWithErrorResponse{
			Error:   true,
			Message: "Failed to create JWT token",
			Data:    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.LogInResponse{
		Error:   false,
		Message: "Success",
		Data: struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: accessToken,
		},
	})
}
