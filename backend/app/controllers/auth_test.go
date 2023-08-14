package controllers

import (
	"api/db/crud"
	"api/db/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testUserUser           *models.User
	testUserAuthentication *models.Authentication
	testUserPassword       *models.Password
)

func testAuthenticationController(t *testing.T) {
	// prepare test user
	testUserUser, _ = crud.AddUser(&models.User{
		Name:      "testerLogIn",
		Authority: models.AuthorityNone,
	})
	testUserAuthentication, _ = crud.AddAuthenticationByUserId(&models.Authentication{
		UserID: testUserUser.ID,
		Email:  "testerLogIn@example.com",
	})
	testUserPassword, _ = crud.AddPasswordByUserId(&models.Password{
		UserID:   testUserUser.ID,
		Password: "testerPw1",
	})

	t.Run("testLogInHandler", testLogInHandler)
}

func testLogInHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.LogInRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Login with existing user -> success",
			method: "POST",
			route:  "/api/v1/auth/login/",
			body: models.LogInRequest{
				Email:    testUserAuthentication.Email,
				Password: testUserPassword.Password,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		//{
		//	description: "login - not existing user",
		//	method:      "POST",
		//	route:       "/api/v1/auth/login/",
		//	body: models.LogInRequest{
		//		Email:    "abc@test.com",
		//		Password: "qpwoeiru",
		//	},
		//	expectedError: true,
		//	expectedCode:  http.StatusBadRequest,
		//},
		//{
		//	description: "login - wrong password format",
		//	method:      "POST",
		//	route:       "/api/v1/auth/login/",
		//	body: models.LogInRequest{
		//		Email:    "cheesecat47@gmail.com",
		//		Password: "pw1",
		//	},
		//	expectedError: true,
		//	expectedCode:  http.StatusBadRequest,
		//},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := json.NewEncoder(buf).Encode(tt.body)
			require.NoError(t, err)
			t.Log("buf: ", buf)

			req := httptest.NewRequest(tt.method, tt.route, buf)
			req.Header.Set("Content-Type", "application/json")
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.LogInResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}
