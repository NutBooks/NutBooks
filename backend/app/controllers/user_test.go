package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Post("/", AddUser)

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
		if err != nil {
			t.Error("Case #", i, ": Failed to convert test case to body param, ", err)
			t.Fail()
		}

		req := httptest.NewRequest(tt.method, tt.route, buf)
		req.Header.Set("Content-Type", "application/json")

		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("Case #", i, ": resp: ", resp)
		assert.NoError(t, err)

		result := &models.AddUserResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			t.Error("Error while parsing response: ", err)
			t.Fail()
		}

		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

func TestGetUserById(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// User
	user := route.Group("/user")
	user.Get("/:id", GetUserById)

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
		if err != nil {
			t.Error("Case #", i, ": Failed to convert test case to body param, ", err)
			t.Fail()
		}

		req := httptest.NewRequest(tt.method, fmt.Sprintf("%s%d", tt.route, tt.body.ID), buf)
		req.Header.Set("Content-Type", "application/json")

		t.Log("Case #", i, ": req: ", req)

		resp, err := app.Test(req, -1)
		t.Log("Case #", i, ": resp: ", resp)
		assert.NoError(t, err)

		result := &models.GetUserByIdResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			t.Error("Error while parsing response: ", err)
			t.Fail()
		}

		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}
