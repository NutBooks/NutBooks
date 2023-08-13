package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func AddBookmark(bookmark *models.Bookmark) (*models.Bookmark, error) {
	log.Debugw("[func AddBookmark]", "bookmark", bookmark)

	result := conn.DB.Create(bookmark)
	if result == nil {
		return nil, errors.New("[func AddBookmark]Failed to create this bookmark")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Debugw("[func AddBookmark]", "bookmark", bookmark)

	return bookmark, nil
}

func GetBookmarkById(id uint) (*models.Bookmark, error) {
	log.Debugw("[func GetBookmarkById]", "id", id)

	found := &models.Bookmark{}
	result := conn.DB.First(found, id)
	if result == nil {
		return nil, errors.New("[func GetBookmarkById] Cannot find this bookmark")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Debugw("[func GetBookmarkById]", "found", found)

	return found, nil
}

func GetAllBookmarks(params *models.GetAllBookmarksRequest) ([]models.Bookmark, error) {
	log.Debugw("[func GetAllBookmarks]", "params", params)

	var found []models.Bookmark
	var result *gorm.DB
	if params.Limit == 0 && params.Offset == 0 {
		result = conn.DB.Find(&found)
	} else {
		result = conn.DB.Limit(params.Limit).Offset(params.Offset).Find(&found)
	}
	if result == nil {
		return nil, errors.New("[func GetAllBookmarks] Failed to find bookmarks")
	}
	log.Debugw("[func GetAllBookmarks]", "found", found)

	return found, nil
}
