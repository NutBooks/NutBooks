package models

import "gorm.io/gorm"

// Bookmark has many Keywords
// https://gorm.io/docs/has_many.html
type Bookmark struct {
	gorm.Model `gorm:"serializer:json"`
	UserID     uint
	Title      string `gorm:"index;"`
	Link       string `gorm:"not null;"`
	//Keywords []Keyword `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type (
	AddBookmarkRequest struct {
		UserID uint   `json:"user_id"`
		Title  string `json:"title"`
		Link   string `json:"link" example:"https://cheesecat47.github.io"`
	}

	AddBookmarkResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetBookmarkByIdRequest struct {
		ID int `json:"id"`
	}

	GetBookmarkByIdResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetAllBookmarksRequest struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}

	GetAllBookmarksResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)
