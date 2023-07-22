package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := Setup()

	t.Helper()

	req := httptest.NewRequest(
		"GET",
		"/api/v1/",
		nil,
	)

	t.Log("req: ", req)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status not OK")

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hi", string(body))
}
