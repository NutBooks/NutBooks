package models

import "gorm.io/gorm"

// 인증(Authorization) 정보 관련 모델 정의

type (
	SocialLogin struct {
		gorm.Model  `gorm:"serializer:json"`
		UserID      uint   `gorm:"not null;unique"`
		SocialCode  uint   `gorm:"not null;"`
		ExternalID  string `gorm:"not null;unique"`
		AccessToken string `gorm:"not null;unique"`
	}

	// Password : 유저 비밀번호 정보 구조체
	Password struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique"`
		Salt       uint

		// 영문 + 숫자 8-12자리
		Password string `gorm:"not null;" json:"password"`
	}

	Cidi struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique"`
		Ci         string
		Di         string
	}
)

// 로그인 핸들러[api/app/controllers.LogInHandler]에서 사용하는 요청/응답 구조체
type (
	LogInRequest struct {
		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:"cheesecat47@gmail.com"`

		// 영문 + 숫자 8-12자리
		Password string `validate:"required,min=8,max=12,alphanum" json:"password" form:"password" example:"qwerty123"`
	}

	LogInResponse struct {
		Error   bool      `json:"error"`
		Data    *Password `json:"data"` // ID가 일치하는 비밀번호 정보
		Message string    `json:"message"`
	}

	LogInWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// 인증 미들웨어에서 사용하는 에러 응답 구조체
type (
	JwtErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)
