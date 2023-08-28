package models

import "gorm.io/gorm"

// 북마크 정보 관련 모델 정의
//
// # References
//   - https://gorm.io/docs/has_many.html

// Bookmark : 북마크 정보 구조체
type Bookmark struct {
	gorm.Model `gorm:"serializer:json"`
	UserID     uint

	// 북마크 제목. 공백이면 og:title로 대체
	Title string `gorm:"index;"`

	// 북마크(웹사이트) 링크. 자세한 형식은 [go-playground/validator] 참고
	//
	// [go-playground/validator]: https://github.com/go-playground/validator
	Link string `gorm:"not null;"`
	//Keywords []Keyword `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// [api/app/controllers.AddBookmarkHandler]에서 사용하는 요청/응답 구조체
type (
	AddBookmarkRequest struct {
		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email"`

		// 북마크 제목. 공백이면 og:title로 대체
		Title string `default:"" json:"title"`

		// 북마크(웹사이트) 링크. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Link string `validate:"required,http_url" json:"link" example:"https://cheesecat47.github.io"`
	}

	AddBookmarkResponse struct {
		Error   bool      `json:"error"`
		Data    *Bookmark `json:"data"` // 생성된 북마크 정보
		Message string    `json:"message"`
	}

	AddBookmarkWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// [api/app/controllers.GetBookmarkByIdHandler]에서 사용하는 요청/응답 구조체
type (
	GetBookmarkByIdRequest struct {
		ID uint `validate:"required,number,min=1" json:"id"`

		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email"`
	}

	GetBookmarkByIdResponse struct {
		Error   bool      `json:"error"`
		Data    *Bookmark `json:"data"` // ID가 일치하는 북마크 정보
		Message string    `json:"message"`
	}

	GetBookmarkByIdWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

// [api/app/controllers.GetAllBookmarksHandler]에서 사용하는 요청/응답 구조체
type (
	GetAllBookmarksRequest struct {
		// 5~50자 길이. 자세한 형식은 [go-playground/validator] 참고
		//
		// [go-playground/validator]: https://github.com/go-playground/validator
		Email string `validate:"required,min=5,max=50,email" json:"email" form:"email"`

		// 특정 id부터 조회할 때 사용
		Offset int `validate:"number,min=0" json:"offset"`

		// limit 개수만큼 조회할 때 사용
		Limit int `validate:"number,min=0" json:"limit"`
	}

	GetAllBookmarksResponse struct {
		Error bool `json:"error"`
		Data  struct {
			Data []Bookmark `json:"data"` // 북마크 정보 배열
			Size int        `json:"size"` // 북마크 배열 길이
		} `json:"data"`
		Message string `json:"message"`
	}

	GetAllBookmarksWithErrorResponse struct {
		Error   bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)
