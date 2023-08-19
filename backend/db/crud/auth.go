package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddAuthenticationByUserId
// UserId를 사용해 Authentication 테이블에 레코드 추가
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
		return nil, errors.New("Failed to create this authentication")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddAuthenticationByUserId]", "authentication", authentication)

	return authentication, nil
}

// AddPasswordByUserId
//
// # UserId를 사용해 Password 테이블에 레코드 추가
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
		return nil, errors.New("Failed to create this password")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddPasswordByUserId]", "password", password)

	return password, nil
}

func GetPasswordByUserId(id uint) (*models.Password, error) {
	found := &models.Password{}
	result := conn.DB.Where("user_id = ?", id).First(found)
	if result == nil {
		return nil, errors.New("Cannot find this password")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}
