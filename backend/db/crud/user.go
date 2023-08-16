package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddUser
// DB의 User 테이블에 레코드 추가
//
// # Parameters
//   - user *models.User: 생성할 유저 정보 구조체
//
// # Returns
//   - *models.User: 생성된 유저 정보 구조체
//   - error: Custom error or [gorm.DB.Error]
func AddUser(user *models.User) (*models.User, error) {
	log.Debugw("[func AddUser]", "user", user)

	result := conn.DB.Create(user)
	if result == nil {
		return nil, errors.New("[func AddUser] Failed to create this user")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddUser]", "user", user)

	return user, nil
}

func GetUserById(id uint) (*models.User, error) {
	log.Debugw("[func GetUserById]", "id", id)

	found := &models.User{}
	result := conn.DB.First(found, id)
	if result == nil {
		return nil, errors.New("[func GetUserById] Cannot find this user")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetUserById]", "found", found)

	return found, nil
}

func GetAllUsers(params *models.GetAllUsersRequest) ([]models.User, error) {
	log.Debugw("[func GetAllUsers]", "params", params)

	var found []models.User
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = conn.DB.Find(&found)
	} else {
		result = conn.DB.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}
	if result == nil {
		return nil, errors.New("[func GetAllUsers] Failed to find users")
	}
	log.Debugw("[func GetAllUsers]", "found", found)

	return found, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	log.Debugw("[func GetUserByEmail]", "email", email)

	found := &models.Authentication{}
	result := conn.DB.Where("email = ?", email).First(&found)
	if result == nil {
		return nil, errors.New("[func GetUserByEmail] Cannot find this user")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetUserByEmail]", "found", found)

	user := &models.User{}
	result = conn.DB.First(user, found.UserID)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetUserByEmail]", "user", user)

	return user, nil
}
