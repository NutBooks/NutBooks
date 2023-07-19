package models

import "gorm.io/gorm"

type SocialLogin struct {
	gorm.Model
	UserID      uint   `gorm:"not null;unique"`
	SocialCode  uint   `gorm:"not null;"`
	ExternalID  string `gorm:"not null;unique"`
	AccessToken string `gorm:"not null;unique"`
}

type Password struct {
	UserID   uint `gorm:"not null;unique"`
	Salt     uint
	Password string `gorm:"not null;"`
}

type Cidi struct {
	UserID uint `gorm:"not null;unique"`
	Ci     string
	Di     string
}
