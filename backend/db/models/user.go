package models

import "gorm.io/gorm"

// 유저 정보 관련 모델 정의
//
// # References
//   - https://rastalion.me/회원-가입-및-로그인을-위한-테이블-설계/
//   - https://gorm.io/docs/has_many.html

type (
	// User : 유저 정보 구조체
	User struct {
		gorm.Model `gorm:"serializer:json"`

		// 알파벳, 숫자, 유니코드 문자 사용 가능
		Name string `gorm:"not null;" json:"name"`

		// 사용자 권한
		//  - [AuthorityNone]
		//  - [AuthorityAdmin]
		Authority string `gorm:"not null;" json:"Authority"`
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

	// Authentication : 유저 이메일 정보 구조체
	Authentication struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint `gorm:"not null;unique;"`

		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `gorm:"not null;unique;" json:"email"`
	}
)

// 회원 가입 핸들러[api/app/controllers.AddUserHandler]에서 사용하는 요청/응답 구조체
type (
	AddUserRequest struct {
		// 알파벳, 숫자, 유니코드 문자 사용 가능
		Name string `validate:"required,min=1,max=50,alphanumunicode" json:"name" form:"name"`

		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:"cheesecat47@gmail.com"`

		// 영문 + 숫자 8-12자리
		Password string `validate:"required,min=8,max=12,alphanum" json:"password" form:"password" example:"qwerty123"`
	}

	AddUserResponse struct {
		Error   bool   `json:"error"`
		Data    *User  `json:"data"` // 생성된 유저 정보
		Message string `json:"message"`
	}

	AddUserWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// [api/app/controllers.GetUserByIdHandler]에서 사용하는 요청/응답 구조체
type (
	GetUserByIdRequest struct {
		ID uint `validate:"required,number,min=1" json:"id"`
	}

	GetUserByIdResponse struct {
		Error   bool   `json:"error"`
		Data    *User  `json:"data"` // ID가 일치하는 유저 정보
		Message string `json:"message"`
	}

	GetUserByIdWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// [api/app/controllers.GetAllUsersHandler]에서 사용하는 요청/응답 구조체
type (
	GetAllUsersRequest struct {
		// 특정 id부터 조회할 때 사용
		Offset int `validate:"number,min=0" json:"offset"`

		// limit 개수만큼 조회할 때 사용
		Limit int `validate:"number,min=0" json:"limit"`
	}

	GetAllUsersResponse struct {
		Error bool `json:"error"`
		Data  struct {
			Data []User `json:"data"` // 유저 정보 배열
			Size int    `json:"size"` // 유저 배열 길이
		} `json:"data"`
		Message string `json:"message"`
	}

	GetAllUsersWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// 이메일 중복 체크 핸들러 [api/app/controllers.CheckEmailDuplicateHandler]에서 사용하는 요청/응답 구조체
type (
	CheckEmailDuplicateRequest struct {
		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:"cheesecat47@gmail.com"`
	}

	CheckEmailDuplicateResponse struct {
		Error   bool            `json:"error"`
		Data    *Authentication `json:"data"` // 생성된 인증 정보
		Message string          `json:"message"`
	}

	CheckEmailDuplicateWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

const (
	AuthorityNone  string = "None"
	AuthorityAdmin string = "Admin"
)
