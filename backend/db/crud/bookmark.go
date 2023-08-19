package crud

import (
	conn "api/db"
	"api/db/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// AddBookmark : DB의 Bookmark 테이블에 레코드 추가
//
// # Parameters
//   - bookmark *models.Bookmark: 생성할 북마크 정보 구조체
//
// # Returns
//   - *models.Bookmark: 생성된 북마크 정보 구조체, 생성 실패 시 nil
//   - error: Custom error or [gorm.DB.Error]
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

// GetBookmarkById : ID가 일치하는 북마크 정보 중 첫 번째 정보를 반환
//
// # Parameters
//   - id uint: 북마크 ID
//
// # Returns
//   - *models.Bookmark: DB에 id가 존재하면 북마크 객체 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
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

// GetAllBookmarks : [offset:offset+limit] 범위에 해당하는 북마크 정보 배열 반환. offset, limit 입력 없을 시 전체 목록 반환.
//
// # Parameters
//   - params *models.GetAllBookmarksRequest: offset(int), limit(int) 포함
//
// # Returns
//   - []models.Bookmark: DB에 해당 범위가 존재하면 북마크 객체 목록 반환, 없다면 nil 반환
//   - error: Custom error or [gorm.DB.Error]
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
