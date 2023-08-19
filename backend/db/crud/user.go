package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddUser : DB의 User 테이블에 레코드 추가
//
// # Parameters
//   - user *models.User: 생성할 유저 정보 구조체
//   - tx *gorm.DB: 트랜잭션 필요 시 사용. nil이면 기본 커넥션 [../db.DB] 객체 사용
//
// # Returns
//   - *models.User: 생성된 유저 정보 구조체, 생성 실패 시 nil
//   - error: Custom error or [gorm.DB.Error]
func AddUser(user *models.User, tx *gorm.DB) (*models.User, error) {
	log.Debugw("[func AddUser]", "user", user)

	if tx == nil {
		tx = conn.DB
	}

	result := tx.Create(user)
	if result == nil {
		return nil, errors.New("[func AddUser] Failed to create this user")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddUser]", "user", user)

	return user, nil
}

// GetUserById : ID가 일치하는 유저 정보 중 첫 번째 정보를 반환
//
// # Parameters
//   - id uint: 유저 ID
//
// # Returns
//   - *models.User: DB에 id가 존재하면 유저 객체 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
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

// GetAllUsers : [offset:offset+limit] 범위에 해당하는 유저 정보 배열 반환. offset, limit 입력 없을 시 전체 목록 반환.
//
// # Parameters
//   - params *models.GetAllUsersRequest: offset(int), limit(int) 포함
//
// # Returns
//   - []models.User: DB에 해당 범위가 존재하면 유저 객체 목록 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
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

// GetUserByEmail : Email이 일치하는 유저 정보 중 첫 번째 정보를 반환
//
// # Parameters
//   - email string: 유저 Email
//
// # Returns
//   - *models.User: DB에 email이 존재하면 유저 객체 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
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
		return nil, errors.New("[func GetUserByEmail] Cannot find this user")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetUserByEmail]", "user", user)

	return user, nil
}
