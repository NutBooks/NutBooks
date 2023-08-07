package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
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
