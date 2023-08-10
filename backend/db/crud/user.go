package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"gorm.io/gorm"
)

func AddUser(user *models.User) (*models.User, error) {
	result := conn.DB.Create(user)
	if result == nil {
		return nil, errors.New("Failed to create this user")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserById(id uint) (*models.User, error) {
	found := &models.User{}
	result := conn.DB.First(found, id)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}

func GetAllUsers(params *models.GetAllUsersRequest) ([]models.User, error) {
	var found []models.User
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = conn.DB.Find(&found)
	} else {
		result = conn.DB.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}

	if result == nil {
		return nil, errors.New("Failed to find users")
	}

	return found, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	found := &models.Authentication{}
	result := conn.DB.Where("email = ?", email).First(&found)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	user := &models.User{}
	result = conn.DB.First(user, found.UserID)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return user, nil
}
