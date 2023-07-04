package models

import "gorm.io/gorm"

type Keyword struct {
	gorm.Model
	//Keyword    string `gorm:"not null;index;"`
	BookmarkID uint
}
