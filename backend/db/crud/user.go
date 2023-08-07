package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"gorm.io/gorm"
)

func AddUser(user *models.User) (*models.User, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Create(user)
	if result == nil {
		return nil, errors.New("Failed to create this user")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserById(id uint) (*models.User, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	found := &models.User{}
	result := db.First(found, id)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}

func GetAllUsers(params *models.GetAllUsersRequest) ([]models.User, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	var found []models.User
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = db.Find(&found)
	} else {
		result = db.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}

	if result == nil {
		return nil, errors.New("Failed to find users")
	}

	return found, nil
}

func GetUserIdByEmail(email string) (*models.User, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	found := &models.User{}
	result := db.Where("email = ?", email).First(&found)
	if result == nil {
		return nil, errors.New("Cannot find this user")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}
