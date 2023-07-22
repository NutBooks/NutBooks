package controllers

import (
	"api/app/utils"
	conn "api/db"
	"api/db/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = validator.New()

func ValidateAddUserRequestParams(p models.AddUserRequest) []*utils.ErrorResponse {
	var errors []*utils.ErrorResponse
	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e utils.ErrorResponse
			e.FailedField = err.StructNamespace()
			e.Tag = err.Tag()
			e.Value = err.Param()
			errors = append(errors, &e)
		}
	}
	return errors
}

// AddUser
//
//	@Summary	새 유저를 추가하는 API
//	@Tags		user
//	@Produce	json
//	@Param		params	body		models.AddUserRequest	true	"body params"
//	@Success	200		{object}	models.AddUserResponse{data=models.User}
//	@Failure	400		{object}	models.AddUserResponse{}
//	@Failure	500		{object}	models.AddUserResponse{}
//	@BasePath	/api/v1
//	@Router		/api/v1/user/ [post]
func AddUser(c *fiber.Ctx) error {
	params := &models.AddUserRequest{}

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if params.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "Name is required",
			Data:    nil,
		})
	}

	if params.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "Email is required",
			Data:    nil,
		})
	}

	validateErr := ValidateAddUserRequestParams(*params)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AddUserResponse{
			Error:   true,
			Message: "Wrong email format",
			Data:    nil,
		})
	}

	user := models.User{
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

	if err = db.Create(&user).Error; err != nil {
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
//	@BasePath	/api/v1
//	@Router		/api/v1/user/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	param := models.GetUserByIdRequest{}

	if err := c.ParamsParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetUserByIdResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
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

	var found models.User
	result := db.First(&found, param.ID)
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
