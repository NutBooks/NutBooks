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
