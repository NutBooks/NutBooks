package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

func TestAddBookmark(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")
	route.Post("/bookmark/new", AddBookmark)

	t.Helper()

	testCases := []models.Bookmark{
		{ // case 0: without userId
			Title: "bookmark_test case 0 title",
			Link:  "https://cheesecat47.github.io/bookmark_test/case0/link",
		},
		{ // case 1: without title
			UserID: 0,
			Link:   "https://cheesecat47.github.io/bookmark_test/case1/link",
		},
		{ // case 2: without link - this should return error
			UserID: 0,
			Title:  "bookmark_test case 2 title",
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(tt)
		if err != nil {
			t.Log("Case #", i, ": Failed to convert test case to body param, ", err)
			continue
		}

		req := httptest.NewRequest(
			"POST",
			"/api/v1/bookmark/new",
			&buf,
		)
		req.Header.Set("Content-Type", "application/json")

		t.Log("req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		t.Log("err: ", err)
		assert.NoError(t, err)
		if i == 2 {
			assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		} else {
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		}
	}
}

func TestGetBookmarkById(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")
	route.Get("/bookmark/:id", GetBookmarkById)

	t.Helper()

	testCases := []string{
		// case 0: get bookmark which index is 1. success
		"1",
		// case 1: get bookmark which index is 0, fail
		"0",
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		req := httptest.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/bookmark/%s", tt),
			nil,
		)

		t.Log("req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		t.Log("err: ", err)
		if i == 1 {
			var result GetBookmarkByIdJSONResult
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Log("err while parsing response:")
			}
			assert.Equal(t, gorm.ErrRecordNotFound.Error(), result.Message)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		}
	}
}

func TestGetAllBookmarks(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")
	route.Get("/bookmark/", GetAllBookmarks)

	t.Helper()

	testCases := []GetAllBookmarksQueryParams{
		// case 0: get all bookmark
		{},
		// case 1: get bookmarks using limit and offset
		{
			Offset: 2,
			Limit:  2,
		},
		// case 2: no offset -> error
		{
			Limit: 2,
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		req := httptest.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/bookmark?offset=%d&limit=%d", tt.Offset, tt.Limit),
			nil,
		)

		t.Log("req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		t.Log("err: ", err)
		if i == 1 {
			var result GetAllBookmarksJSONResult
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Log("err while parsing response:")
			}
			assert.Equal(t, 2, len(result.Data))
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		} else if i == 2 {
			var result GetAllBookmarksJSONResult
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Log("err while parsing response:")
			}
			assert.Equal(t, "limit, offset은 같이 설정되어야 하고, 0 이상 정수를 입력해주세요.", result.Message)
			assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		}
	}
}
