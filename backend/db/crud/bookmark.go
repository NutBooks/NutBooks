package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"gorm.io/gorm"
)

func AddBookmark(bookmark *models.Bookmark) error {
	db, err := conn.GetDB()
	if err != nil {
		return err
	}

	result := db.Create(bookmark)
	if result == nil {
		return errors.New("Failed to create this bookmark")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetBookmarkById(id uint) (*models.Bookmark, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	found := &models.Bookmark{}
	result := db.First(found, id)
	if result == nil {
		return nil, errors.New("Cannot find this bookmark")
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return found, nil
}

func GetAllBookmarks(params *models.GetAllBookmarksRequest) ([]models.Bookmark, error) {
	db, err := conn.GetDB()
	if err != nil {
		return nil, err
	}

	var found []models.Bookmark
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = db.Find(&found)
	} else {
		result = db.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}

	if result == nil {
		return nil, errors.New("Failed to find bookmarks")
	}

	return found, nil
}
