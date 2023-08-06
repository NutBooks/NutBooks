package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddBookmark(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// bookmark
	bookmark := route.Group("/bookmark")
	bookmark.Post("/", AddBookmarkHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.AddBookmarkRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description: "case 0: without userId",
			method:      "POST",
			route:       "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				Title: "bookmark_test case 0 title",
				Link:  "https://cheesecat47.github.io/bookmark_test/case0/link",
			},
			expectedError: false,
			expectedCode:  http.StatusOK,
			expectedBody:  "Success",
		},
		{
			description: "case 1: without title",
			method:      "POST",
			route:       "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				UserID: 0,
				Link:   "https://cheesecat47.github.io/bookmark_test/case1/link",
			},
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
		{
			description: "case 2: without link - this should return error",
			method:      "POST",
			route:       "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				UserID: 0,
				Title:  "bookmark_test case 2 title",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "link is required parameter",
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(tt.body)
		require.NoError(t, err)

		req := httptest.NewRequest(tt.method, tt.route, buf)
		req.Header.Set("Content-Type", "application/json")
		t.Log("req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

func TestGetBookmarkById(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// bookmark
	bookmark := route.Group("/bookmark")
	bookmark.Get("/:id/", GetBookmarkByIdHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.GetBookmarkByIdRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "case 0: get bookmark which index is 1. success",
			method:        "GET",
			route:         "/api/v1/bookmark/1/",
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
		{
			description:   "case 1: get bookmark which index is 0, fail",
			method:        "GET",
			route:         "/api/v1/bookmark/0/",
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		req := httptest.NewRequest(tt.method, tt.route, nil)
		req.Header.Set("Content-Type", "application/json")
		t.Log("req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

func TestGetAllBookmarks(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// bookmark
	bookmark := route.Group("/bookmark")
	bookmark.Get("/", GetAllBookmarksHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.GetAllBookmarksRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "case 0: get all bookmark",
			method:        "GET",
			route:         "/api/v1/bookmark/",
			body:          models.GetAllBookmarksRequest{},
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
		{
			description: "case 1: get bookmarks using limit and offset",
			method:      "GET",
			route:       "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Offset: 0,
				Limit:  1,
			},
			expectedError: false,
			expectedCode:  http.StatusBadRequest,
		},
		{
			description: "case 2: no offset -> error",
			method:      "GET",
			route:       "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Limit: 2,
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		req := httptest.NewRequest(
			tt.method,
			fmt.Sprintf("%s?offset=%d&limit=%d", tt.route, tt.body.Offset, tt.body.Limit),
			nil)
		req.Header.Set("Content-Type", "application/json")
		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("resp: ", resp)
		require.NoError(t, err)

		result := &models.GetAllUsersResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}
