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

// GetAllUsers
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
func GetAllUsers(c *fiber.Ctx) error {
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

	db, err := conn.GetDB()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var found []models.User
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = db.Find(&found)
	} else {
		result = db.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}

	if result == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetAllUsersResponse{
			Error:   true,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GetAllUsersResponse{
		Error:   false,
		Message: "Success",
		Data:    found,
	})
}
