package models

import "gorm.io/gorm"

type (
	SocialLogin struct {
		gorm.Model  `gorm:"serializer:json"`
		UserID      uint   `gorm:"not null;unique"`
		SocialCode  uint   `gorm:"not null;"`
		ExternalID  string `gorm:"not null;unique"`
		AccessToken string `gorm:"not null;unique"`
	}

	Password struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique"`
		Salt       uint
		Password   string `validate:"min=8,max=12,alphanum" gorm:"not null;"`
	}

	Cidi struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique"`
		Ci         string
		Di         string
	}
)

type (
	LogInRequest struct {
		Email    string `validate:"required,min=5,max=50,email" json:"email" example:""`
		Password string `validate:"required,min=8,max=12,alphanum" json:"password" example:"비밀번호는 영문 + 숫자 8-12자리"`
	}

	LogInResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)
