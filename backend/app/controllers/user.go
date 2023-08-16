package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddUserHandler
//
//	@Summary	새 유저 추가
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		params	body		models.AddUserRequest	true	"비밀번호는 영문 + 숫자 8-12자리"
//	@Success	201		{object}	models.AddUserResponse{data=models.User}
//	@Failure	400		{object}	models.AddUserResponse{}
//	@Failure	500		{object}	models.AddUserResponse{}
//	@Router		/api/v1/user [post]
func AddUserHandler(c *fiber.Ctx) error {
	params := &models.AddUserRequest{}
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: fmt.Sprintf("Failed to parse parameters: %v", err.Error()),
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
	log.Infow("[func AddUserHandler]", "params", params)

	checkEmail, err := crud.GetUserByEmail(params.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: fmt.Sprintf("Failed to check email duplicate: %v", err.Error()),
			Data:    err,
		})
	}
	if checkEmail != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "This email is already in use",
			Data:    nil,
		})
	}

	user := &models.User{
		Name:      params.Name,
		Authority: models.AuthorityNone,
	}
	user, err = crud.AddUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: fmt.Sprintf("Failed to create user: %v", err.Error()),
			Data:    nil,
		})
	}
	log.Debugw("[func AddUserHandler]", "user", user)

	authentication := &models.Authentication{
		UserID: user.ID,
		Email:  params.Email,
	}
	authentication, err = crud.AddAuthenticationByUserId(authentication)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: fmt.Sprintf("Failed to create authentication: %v", err.Error()),
			Data:    nil,
		})
	}
	log.Debugw("[func AddUserHandler]", "authentication", authentication)

	password := &models.Password{
		UserID:   user.ID,
		Password: params.Password,
	}
	password, err = crud.AddPasswordByUserId(password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}
	log.Debugw("[func AddUserHandler]", "password", password)

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
//	@Success	200	{object}	models.GetUserByIdResponse{data=models.User}
//	@Failure	400	{object}	models.GetUserByIdResponse{}
//	@Failure	500	{object}	models.AddUserResponse{}
//	@Router		/api/v1/user/{id} [get]
func GetUserByIdHandler(c *fiber.Ctx) error {
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
	log.Infow("[func GetUserByIdHandler]", "params", params)

	found, err := crud.GetUserById(params.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
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
//	@Param		offset	query		int	false	"limit과 offset은 같이 입력해야 합니다."
//	@Param		limit	query		int	false	"limit과 offset은 같이 입력해야 합니다."
//	@Success	200		{object}	models.GetAllUsersResponse{data=[]models.User{}}
//	@Failure	400		{object}	models.GetAllUsersResponse{}
//
//	@Failure	500		{object}	models.GetAllUsersResponse{}
//	@Router		/api/v1/user [get]
func GetAllUsersHandler(c *fiber.Ctx) error {
	params := &models.GetAllUsersRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersResponse{
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
	log.Infow("[func GetAllUsersHandler]", "params", params)

	found, err := crud.GetAllUsers(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}
	log.Debugw("[func GetAllUsersHandler]", "found", found)

	return c.Status(fiber.StatusOK).JSON(models.GetAllUsersResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}

// CheckEmailDuplicateHandler
//
//	@Summary		이메일 중복 체크.
//	@Description	입력한 이메일을 사용하는 유저가 있다면 Body의 Message로 True 반환, 없다면 False 반환.
//	@Tags			user
//	@Produce		json
//	@Param			email	query		string	true	"중복 체크 할 이메일 주소 입력. 최대 길이 50자 제한."
//	@Success		200		{object}	models.CheckEmailDuplicateResponse{}
//	@Failure		400		{object}	models.CheckEmailDuplicateResponse{}
//	@Failure		500		{object}	models.CheckEmailDuplicateResponse{}
//	@Router			/api/v1/user/check-email [get]
func CheckEmailDuplicateHandler(c *fiber.Ctx) error {
	params := &models.CheckEmailDuplicateRequest{}
	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CheckEmailDuplicateResponse{
			Error:   true,
			Message: fmt.Sprintf("Failed to parse query parameters: %v", err.Error()),
			Data:    nil,
		})
	}

	validator := &utils.Validator{}
	validateErrs := validator.Validate(params)
	if validateErrs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CheckEmailDuplicateResponse{
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
			return c.Status(fiber.StatusInternalServerError).JSON(models.CheckEmailDuplicateResponse{
				Error:   true,
				Message: err.Error(),
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
