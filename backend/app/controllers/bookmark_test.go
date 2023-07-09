package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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
