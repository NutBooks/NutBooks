package controllers

import (
	"api/configs"
	"api/db/models"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogInHandler(t *testing.T) {
	config := configs.FiberConfig()
	app := fiber.New(config)
	route := app.Group("/api/v1")

	// auth
	auth := route.Group("/auth")
	auth.Post("/login/", LogInHandler)

	t.Helper()

	testCases := []struct {
		description   string
		method        string
		route         string
		body          models.LogInRequest
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description: "login success",
			method:      "POST",
			route:       "/api/v1/auth/login/",
			body: models.LogInRequest{
				Email:    "cheesecat47@gmail.com",
				Password: "qpwoeiru",
			},
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
		{
			description: "login - not existing user",
			method:      "POST",
			route:       "/api/v1/auth/login/",
			body: models.LogInRequest{
				Email:    "abc@test.com",
				Password: "qpwoeiru",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			description: "login - wrong password format",
			method:      "POST",
			route:       "/api/v1/auth/login/",
			body: models.LogInRequest{
				Email:    "cheesecat47@gmail.com",
				Password: "pw1",
			},
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
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

		result := &models.LogInResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		require.NoError(t, err)

		require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
	}
}
