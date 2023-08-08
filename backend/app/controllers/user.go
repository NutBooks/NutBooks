package controllers

import (
	"api/app/utils"
	"api/db/crud"
	"api/db/models"
	"github.com/gofiber/fiber/v2"
)

// AddUserHandler
//
//	@Summary	새 유저를 추가하는 API
//	@Tags		user
//	@Produce	json
//	@Param		params	body		models.AddUserRequest	true	"params"
//	@Success	200		{object}	models.AddUserResponse{data=models.User}
//	@Failure	400		{object}	models.AddUserResponse{}
//	@Failure	500		{object}	models.AddUserResponse{}
//	@Router		/api/v1/user/ [post]
func AddUserHandler(c *fiber.Ctx) error {
	params := &models.AddUserRequest{}
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
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "Validation failed",
			Data:    validateErrs,
		})
	}

	// 이메일 중복 확인 로직을 추가하든, 이메일 중복 시 user create를 rollback 하든
	// 로릭 추가 필요

	user := &models.User{
		Name:      params.Name,
		Authority: models.AuthorityNone,
	}

	user, err := crud.AddUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	authentication := &models.Authentication{
		UserID: user.ID,
		Email:  params.Email,
	}

	authentication, err = crud.AddAuthentication(authentication)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

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

	return c.Status(fiber.StatusOK).JSON(models.AddUserResponse{
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
//	@Router		/api/v1/user/{id}/ [get]
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

	found, err := crud.GetUserById(params.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

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
//	@Router		/api/v1/user/ [get]
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

	found, err := crud.GetAllUsers(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetAllUsersResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
