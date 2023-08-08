package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"gorm.io/gorm"
)

func AddAuthentication(authentication *models.Authentication) (*models.Authentication, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Create(authentication)
	if result == nil {
		return nil, errors.New("Failed to create this authentication")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return authentication, nil
}

func AddPasswordByUserId(password *models.Password) (*models.Password, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Create(password)
	if result == nil {
		return nil, errors.New("Failed to create this password")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return password, nil
}

func GetPasswordByUserId(id uint) (*models.Password, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	found := &models.Password{}
	result := db.Where("user_id = ?", id).First(found)
	if result == nil {
		return nil, errors.New("Cannot find this password")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}
