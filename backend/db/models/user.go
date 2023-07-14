package models

import (
	"gorm.io/gorm"
)

// User has many bookmarks
// https://rastalion.me/회원-가입-및-로그인을-위한-테이블-설계/
// https://gorm.io/docs/has_many.html
type User struct {
	gorm.Model
	Name      string `gorm:"not null;"`
	Authority string `gorm:"not null;"`
}

type Profile struct {
	gorm.Model
	UserID   uint `gorm:"not null;unique;"`
	Nickname string
	ImageURL string
}

type Authority struct {
	gorm.Model
	UserID    uint   `gorm:"not null;unique;"`
	Authority string `gorm:"not null;"`
}

type Authentication struct {
	UserID uint   `gorm:"not null;unique;"`
	Email  string `gorm:"not null;unique;"`
}

//type Subscription struct {
//
//}
