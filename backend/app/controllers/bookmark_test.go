package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBookmarkWithoutUserId(t *testing.T) {
	testCases := []struct {
		userId   uint
		title    string
		link     string
		keywords []string
	}{
		{ // case 1: without userId
			userId:   0,
			title:    "",
			link:     "https://github.com/TangoEnSkai/uber-go-style-guide-kr#테스트-테이블-test-tables",
			keywords: []string{""},
		},
	}

	for _, tt := range testCases {
		_, err := AddBookmark(tt.userId, tt.title, tt.link, tt.keywords)
		assert.EqualError(t, err, "No user id")
	}
}
