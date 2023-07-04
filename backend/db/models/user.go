package models

import (
	"gorm.io/gorm"
)

// User has many bookmarks
// https://gorm.io/docs/has_many.html
type User struct {
	gorm.Model
	GitHubID    string `gorm:"not null;unique;"`
	Password    string `gorm:"not null;"`
	Name        string `gorm:"not null;"`
	AccessToken string
	//Bookmarks   []Bookmark `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
