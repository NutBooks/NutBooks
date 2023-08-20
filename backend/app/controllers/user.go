package controllers

import (
	"api/app/middlewares"
	"api/app/utils"
	conn "api/db"
	"api/db/crud"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// AddUserHandler
//
//	@Summary	새 유저 추가
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		params	body		models.AddUserRequest	true	"비밀번호는 영문 + 숫자 8-12자리"
//	@Success	201		{object}	models.AddUserResponse
//	@Failure	400		{object}	models.AddUserWithErrorResponse
//	@Failure	500		{object}	models.AddUserWithErrorResponse
//	@Router		/api/v1/user [post]
func AddUserHandler(c *fiber.Ctx) error {
	params := &models.AddUserRequest{}
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func AddUserHandler]", "params", params)

	checkEmail, err := crud.GetUserByEmail(params.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserWithErrorResponse{
				Error:   true,
				Message: "Failed to check email duplicate",
				Data:    err,
			})
		}
	}
	if checkEmail != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "This email is already in use",
			Data:    nil,
		})
	}

	tx := conn.DB.Begin()

	user := &models.User{
		Name:      params.Name,
		Authority: models.AuthorityNone,
	}
	user, err = crud.AddUser(user, tx)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "Failed to create user",
			Data:    err,
		})
	}
	log.Debugw("[func AddUserHandler]", "user", user)

	authentication := &models.Authentication{
		UserID: user.ID,
		Email:  params.Email,
	}
	authentication, err = crud.AddAuthenticationByUserId(authentication, tx)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "Failed to create authentication",
			Data:    err,
		})
	}
	log.Debugw("[func AddUserHandler]", "authentication", authentication)

	password := &models.Password{
		UserID:   user.ID,
		Password: params.Password,
	}
	password, err = crud.AddPasswordByUserId(password, tx)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserWithErrorResponse{
			Error:   true,
			Message: "Failed to create password",
			Data:    err,
		})
	}
	log.Debugw("[func AddUserHandler]", "password", password)

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(models.AddUserResponse{
		Error:   false,
		Message: "Success",
		Data:    user,
	})
}

// GetUserByIdHandler
//
//	@Summary	UserID를 사용해 유저 1명 정보 읽기
//	@Tags		user
//	@Produce	json
//	@Param		id	path		uint	true	"User ID"
//	@Success	200	{object}	models.GetUserByIdResponse
//	@Failure	400	{object}	models.GetUserByIdWithErrorResponse
//	@Failure	401	{object}	models.GetUserByIdWithErrorResponse
//	@Failure	500	{object}	models.GetUserByIdWithErrorResponse
//	@Router		/api/v1/user/{id} [get]
func GetUserByIdHandler(c *fiber.Ctx) error {
	params := &models.GetUserByIdRequest{}
	err := c.ParamsParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func GetUserByIdHandler]", "params", params)

	token := c.Locals("user_id").(*jwt.Token)
	err = middlewares.ValidToken(token, params.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.GetUserByIdWithErrorResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	found, err := crud.GetUserById(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdWithErrorResponse{
				Error:   true,
				Message: "record not found",
				Data:    err,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.GetUserByIdWithErrorResponse{
				Error:   true,
				Message: "Failed to get user",
				Data:    err,
			})
		}
	}
	log.Debugw("[func GetUserByIdHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetUserByIdResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}

// GetAllUsersHandler
//
//	@Summary	모든 유저 목록 반환
//	@Tags		user
//	@Produce	json
//	@Param		offset	query		int	false	"특정 id부터 조회할 때 사용"
//	@Param		limit	query		int	false	"limit 개수만큼 조회할 때 사용"
//	@Success	200		{object}	models.GetAllUsersResponse
//	@Failure	400		{object}	models.GetAllUsersWithErrorResponse
//	@Failure	401		{object}	models.GetAllUsersWithErrorResponse
//	@Failure	500		{object}	models.GetAllUsersWithErrorResponse
//	@Router		/api/v1/user [get]
func GetAllUsersHandler(c *fiber.Ctx) error {
	// check authentication

	params := &models.GetAllUsersRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersWithErrorResponse{
			Error:   true,
			Message: "Failed to parse parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func GetAllUsersHandler]", "params", params)

	token := c.Locals("user_id").(*jwt.Token)
	err = middlewares.ValidToken(token, params.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.GetAllUsersWithErrorResponse{
			Error:   true,
			Message: err.Error(),
			Data:    err,
		})
	}

	found, err := crud.GetAllUsers(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersWithErrorResponse{
			Error:   true,
			Message: "Failed to get users",
			Data:    err,
		})
	}
	log.Debugw("[func GetAllUsersHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetAllUsersResponse{
		Error:   false,
		Message: "Success",
		Data: struct {
			Data []models.User `json:"data"`
			Size int           `json:"size"`
		}{
			Data: found,
			Size: len(found),
		},
	})
}

// CheckEmailDuplicateHandler
//
//	@Summary		이메일 중복 체크.
//	@Description	입력한 이메일을 사용하는 유저가 있다면 Body의 Message로 True 반환, 없다면 False 반환.
//	@Tags			user
//	@Produce		json
//	@Param			email	query		string	true	"중복 체크 할 이메일 주소 입력. 최대 길이 50자 제한."
//	@Success		200		{object}	models.CheckEmailDuplicateResponse
//	@Failure		400		{object}	models.CheckEmailDuplicateWithErrorResponse
//	@Failure		500		{object}	models.CheckEmailDuplicateWithErrorResponse
//	@Router			/api/v1/user/check-email [get]
func CheckEmailDuplicateHandler(c *fiber.Ctx) error {
	params := &models.CheckEmailDuplicateRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CheckEmailDuplicateWithErrorResponse{
			Error:   true,
			Message: "Failed to parse query parameters",
			Data:    err,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CheckEmailDuplicateWithErrorResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}
	log.Infow("[func CheckEmailDuplicateHandler]", "params", params)

	_, err = crud.GetUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusOK).JSON(models.CheckEmailDuplicateResponse{
				Error:   false,
				Message: "False",
				Data:    nil,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.CheckEmailDuplicateWithErrorResponse{
				Error:   true,
				Message: "Failed to get email",
				Data:    err,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.CheckEmailDuplicateResponse{
		Error:   false,
		Message: "True",
		Data:    nil,
	})
}
