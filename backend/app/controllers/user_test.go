package controllers

import (
	"api/db/crud"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testUser1  *models.User
	testEmail1 *models.Authentication
)

func testUserController(t *testing.T) {
	// prepare test user
	testUser1, _ = crud.AddUser(&models.User{
		Name:      "testUserController",
		Authority: models.AuthorityNone,
	})
	testEmail1, _ = crud.AddAuthenticationByUserId(&models.Authentication{
		UserID: testUser1.ID,
		Email:  "testUser1@example.com",
	})

	t.Run("testAddUserHandler", testAddUserHandler)
	t.Run("testGetUserByIdHandler", testGetUserByIdHandler)
	t.Run("testGetAllUsersHandler", testGetAllUsersHandler)
	t.Run("testCheckEmailDuplicateHandler", testCheckEmailDuplicateHandler)
}

func testAddUserHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.AddUserRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Add  user -> success",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "tester1",
				Email:    "tester1@example.com",
				Password: "tester1pw1",
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Add user without name and email -> fail",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "",
				Email:    "",
				Password: "tester2pw1",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Add user without name -> fail",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "",
				Email:    "tester3@example.com",
				Password: "tester3pw1",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Add user without email -> fail",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "tester4",
				Email:    "",
				Password: "tester4pw1",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Add user with wrong email format -> fail",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "tester5",
				Email:    "tester5email",
				Password: "tester5pw1",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Add user with wrong password format -> fail",
			method: "POST",
			route:  "/api/v1/user",
			body: models.AddUserRequest{
				Name:     "tester5",
				Email:    "tester5email",
				Password: "pw1",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
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

			result := &models.AddUserResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}

func testGetUserByIdHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.GetUserByIdRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Get user of testUser1 using testUser1.ID -> success",
			method: "GET",
			route:  "/api/v1/user",
			body: models.GetUserByIdRequest{
				ID: testUser1.ID,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Get not existing user -> fail",
			method: "GET",
			route:  "/api/v1/user",
			body: models.GetUserByIdRequest{
				ID: 464749,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "record not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := json.NewEncoder(buf).Encode(tt.body)
			require.NoError(t, err)
			t.Log("buf: ", buf)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("%s/%d", tt.route, tt.body.ID), buf)
			req.Header.Set("Content-Type", "application/json")
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.GetUserByIdResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}

func testGetAllUsersHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.GetAllUsersRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:            "Get all users -> success",
			method:          "GET",
			route:           "/api/v1/user",
			body:            models.GetAllUsersRequest{},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Get all users using negative offset value -> fail",
			method: "GET",
			route:  "/api/v1/user",
			body: models.GetAllUsersRequest{
				Offset: -1,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Get all users using negative limit value -> fail",
			method: "GET",
			route:  "/api/v1/user",
			body: models.GetAllUsersRequest{
				Limit: -1,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Get all users using offset and limit value -> success",
			method: "GET",
			route:  "/api/v1/user",
			body: models.GetAllUsersRequest{
				Offset: int(testUser1.ID),
				Limit:  1,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				tt.method,
				fmt.Sprintf("%s?offset=%d&limit=%d", tt.route, tt.body.Offset, tt.body.Limit),
				nil,
			)
			req.Header.Set("Content-Type", "application/json")
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.GetAllUsersResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}

// 이메일 중복 체크 핸들러 테스트
//
// [controllers.CheckEmailDuplicateHandler]
//
// # Test Cases
//
//   - Case 1: 쿼리 파라미터로 입력한 이메일이 존재하는 경우 응답 Body의 Message로 "True" (이 이메일 사용 불가)
//   - Case 2: 이메일이 존재하지 않는다면 "False"를 반환. (이 이메일 사용 가능)
//   - Case 3: 이상한 형식의 이메일 입력. Validation Error 반환.
func testCheckEmailDuplicateHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		params          map[string]string
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Check with a new unique email -> 200 and 'True'",
			method: "GET",
			route:  "/api/v1/user/check-email",
			params: map[string]string{
				"email": "unique1@example.com",
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "False",
		},
		{
			name:   "Check with an existing email -> 200 but 'False'",
			method: "GET",
			route:  "/api/v1/user/check-email",
			params: map[string]string{
				"email": testEmail1.Email,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "True",
		},
		{
			name:   "Check with an email of wrong format -> 400",
			method: "GET",
			route:  "/api/v1/user/check-email",
			params: map[string]string{
				"email": "cheesecat47_at_github_com",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				tt.method,
				fmt.Sprintf("%s?email=%s", tt.route, tt.params["email"]),
				nil,
			)
			t.Log("req:", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.CheckEmailDuplicateResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}
