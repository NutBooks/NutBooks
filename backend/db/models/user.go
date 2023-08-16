package models

import "gorm.io/gorm"

// 유저 정보 관련 모델 정의
//
// # References
//   - https://rastalion.me/회원-가입-및-로그인을-위한-테이블-설계/
//   - https://gorm.io/docs/has_many.html

type (
	// User
	//
	// 유저 정보 구조체
	User struct {
		gorm.Model `gorm:"serializer:json"`

		// 사용자 이름.
		//
		// 알파벳, 숫자, 유니코드 문자 사용 가능
		//
		// [AddUserRequest.Name] 으로 입력 시 validation
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

	Authentication struct {
		gorm.Model `gorm:"serializer:json"`
		UserID     uint   `gorm:"not null;unique;"`
		Email      string `validate:"min=5,max=50,email" gorm:"not null;unique;" json:"email"`
	}

	//Subscription struct {
	//
	//}
)

// AddUserRequest
//
// 회원 가입 핸들러[api/app/controllers.AddUserHandler]에서 사용하는 요청/응답 구조체
type (
	AddUserRequest struct {
		// 사용자 이름. 알파벳, 숫자, 유니코드 문자 사용 가능.
		Name string `validate:"required,min=1,max=50,alphanumunicode" json:"name" form:"name" example:""`

		// 이메일 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email" example:""`

		// 비밀번호는 영문 + 숫자 8-12자리
		Password string `validate:"required,min=8,max=12,alphanum" json:"password" form:"password" example:""`
	}

	AddUserResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

type (
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
