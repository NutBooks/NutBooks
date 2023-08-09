package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"gorm.io/gorm"
)

func AddAuthentication(authentication *models.Authentication) (*models.Authentication, error) {
	result := conn.DB.Create(authentication)
	if result == nil {
		return nil, errors.New("Failed to create this authentication")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return authentication, nil
}

func AddPasswordByUserId(password *models.Password) (*models.Password, error) {
	result := conn.DB.Create(password)
	if result == nil {
		return nil, errors.New("Failed to create this password")
	}

	if result.Error != nil {
		return nil, result.Error
	}

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
