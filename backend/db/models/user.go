package models

import "gorm.io/gorm"

// User has many bookmarks
// https://rastalion.me/회원-가입-및-로그인을-위한-테이블-설계/
// https://gorm.io/docs/has_many.html
type (
	User struct {
		gorm.Model `gorm:"serializer:json"`
		Name       string `gorm:"not null;" json:"name"`
		Authority  string `gorm:"not null;" json:"Authority"`
	}

	Profile struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique;"`
		Nickname   string
		ImageURL   string
	}

	Authority struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint   `gorm:"not null;unique;"`
		Authority  string `gorm:"not null;"`
	}

	Authentication struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint   `gorm:"not null;unique;"`
		Email      string `gorm:"not null;unique;" json:"email"`
	}

	//Subscription struct {
	//
	//}
)

type (
	AddUserRequest struct {
		Name  string `validate:"required,min=1,max=20" json:"name" example:""`
		Email string `validate:"required,min=5,max=20,email" json:"email" example:""`
	}

	AddUserResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

type (
	GetUserByIdRequest struct {
		ID int `validate:"required,number,min=1" json:"id"`
	}

	GetUserByIdResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

const (
	AuthorityNone  string = "None"
	AuthorityAdmin string = "Admin"
)
