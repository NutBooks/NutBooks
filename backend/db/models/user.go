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
		Email      string `validate:"min=5,max=50,email" gorm:"not null;unique;" json:"email"`
	}

	//Subscription struct {
	//
	//}
)

type (
	AddUserRequest struct {
		Name     string `validate:"required,min=1,max=50" json:"name" form:"name" example:""`
		Email    string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:""`
		Password string `validate:"required,min=8,max=12,alphanum" json:"password" form:"password" example:"비밀번호는 영문 + 숫자 8-12자리"`
	}

	AddUserResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetUserByIdRequest struct {
		ID uint `validate:"required,number,min=1" json:"id"`
	}

	GetUserByIdResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	GetAllUsersRequest struct {
		Offset int `validate:"number,min=0" json:"offset"`
		Limit  int `validate:"number,min=0" json:"limit"`
	}

	GetAllUsersResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	CheckEmailDuplicateRequest struct {
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:""`
	}

	CheckEmailDuplicateResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

const (
	AuthorityNone  string = "None"
	AuthorityAdmin string = "Admin"
)
