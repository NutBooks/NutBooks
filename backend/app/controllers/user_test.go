package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAddUser(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Post("/", AddUserHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.AddUserRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description: "create user",
			method:      "POST",
			route:       "/api/v1/user/",
			body: models.AddUserRequest{
				Name:  "tester1",
				Email: "tester1@example.com",
			},
			expectedError: false,
			expectedCode:  http.StatusOK,
			expectedBody:  "Success",
		},
		{
			description: "create user without name and email -> error",
			method:      "POST",
			route:       "/api/v1/user/",
			body: models.AddUserRequest{
				Name:  "",
				Email: "",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "Name is required",
		},
		{
			description: "create user without name -> error",
			method:      "POST",
			route:       "/api/v1/user/",
			body: models.AddUserRequest{
				Name:  "",
				Email: "tester3@example.com",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "Name is required",
		},
		{
			description: "create user without email -> error",
			method:      "POST",
			route:       "/api/v1/user/",
			body: models.AddUserRequest{
				Name:  "tester4",
				Email: "",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "Email is required",
		},
		{
			description: "create user with wrong email format -> error",
			method:      "POST",
			route:       "/api/v1/user/",
			body: models.AddUserRequest{
				Name:  "tester5",
				Email: "tester5email",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "Wrong email format",
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(tt.body)
		require.NoError(t, err)

		req := httptest.NewRequest(tt.method, tt.route, buf)
		req.Header.Set("Content-Type", "application/json")
		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("Case #", i, ": resp: ", resp)
		require.NoError(t, err)

		result := &models.AddUserResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
	}
}

func TestGetUserById(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Get("/:id/", GetUserByIdHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.GetUserByIdRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description: "Get user",
			method:      "GET",
			route:       "/api/v1/user/",
			body: models.GetUserByIdRequest{
				ID: 1,
			},
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
	}

	for i, tt := range testCases {
		t.Log("Case #", i, ": ", tt)

		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(tt.body)
		require.NoError(t, err)

		req := httptest.NewRequest(tt.method, fmt.Sprintf("%s%d", tt.route, tt.body.ID), buf)
		req.Header.Set("Content-Type", "application/json")
		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("Case #", i, ": resp: ", resp)
		require.NoError(t, err)

		result := &models.GetUserByIdResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
	}
}

func TestGetAllUsers(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Get("/", GetAllUsersHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.GetAllUsersRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Get all users",
			method:        "GET",
			route:         "/api/v1/user/",
			body:          models.GetAllUsersRequest{},
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
		{
			description: "Get all users",
			method:      "GET",
			route:       "/api/v1/user/",
			body: models.GetAllUsersRequest{
				Offset: -1,
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
			nil,
		)
		req.Header.Set("Content-Type", "application/json")
		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("Case #", i, ": resp: ", resp)
		require.NoError(t, err)

		result := &models.GetAllUsersResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
	}
}
