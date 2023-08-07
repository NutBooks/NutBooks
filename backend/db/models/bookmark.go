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
		UserID uint   `validate:"required,number,min=1" json:"user_id"`
		Title  string `validate:"min=1" default:"" json:"title"`
		Link   string `validate:"required,http_url" json:"link" example:"https://cheesecat47.github.io"`
	}

	AddBookmarkResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetBookmarkByIdRequest struct {
		ID int `validate:"required,number,min=1" json:"id"`
	}

	GetBookmarkByIdResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetAllBookmarksRequest struct {
		Offset int `validate:"number,min=0" json:"offset"`
		Limit  int `validate:"number,min=0" json:"limit"`
	}

	GetAllBookmarksResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)
