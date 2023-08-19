package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddAuthenticationByUserId : Authentication 테이블에 레코드 추가
//
// # Parameters
//   - authentication *models.Authentication: 생성할 인증 정보 구조체
//   - tx *gorm.DB: 트랜잭션 필요 시 사용. nil이면 기본 커넥션 [../db.DB] 객체 사용
//
// # Returns
//   - *models.Authentication: 생성된 유저 정보 구조체
//   - error: Custom error or [gorm.DB.Error]
func AddAuthenticationByUserId(authentication *models.Authentication, tx *gorm.DB) (*models.Authentication, error) {
	log.Debugw("[func AddAuthenticationByUserId]", "authentication", authentication)

	if tx == nil {
		tx = conn.DB
	}

	result := tx.Create(authentication)
	if result == nil {
		return nil, errors.New("[func AddAuthenticationByUserId] Failed to create this authentication")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddAuthenticationByUserId]", "authentication", authentication)

	return authentication, nil
}

// AddPasswordByUserId : Password 테이블에 레코드 추가
//
// # Parameters
//   - password *models.Password: 생성할 비밀번호 정보 구조체
//   - tx *gorm.DB: 트랜잭션 필요 시 사용. nil이면 기본 커넥션 [../db.DB] 객체 사용
//
// # Returns
//   - *models.Password: 생성된 유저 정보 구조체
//   - error: Custom error or [gorm.DB.Error]
func AddPasswordByUserId(password *models.Password, tx *gorm.DB) (*models.Password, error) {
	log.Debugw("[func AddPasswordByUserId]", "password", password)

	if tx == nil {
		tx = conn.DB
	}

	result := tx.Create(password)
	if result == nil {
		return nil, errors.New("[func AddPasswordByUserId] Failed to create this password")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddPasswordByUserId]", "password", password)

	return password, nil
}

// GetPasswordByUserId : ID가 일치하는 비밀번호 정보 반환
//
// # Parameters
//   - id uint: 유저 아이디
//
// # Returns
//   - *models.Password: DB에 userId와 일치하는 비밀번호 정보가 있으면 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
func GetPasswordByUserId(userId uint) (*models.Password, error) {
	log.Debugw("[func GetPasswordByUserId]", "userId", userId)

	found := &models.Password{}
	result := conn.DB.Where("user_id = ?", userId).First(found)
	if result == nil {
		return nil, errors.New("[func GetPasswordByUserId] Cannot find this password")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetPasswordByUserId]", "found", found)

	return found, nil
}
