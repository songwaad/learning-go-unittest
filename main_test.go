package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	app := setup()

	// Define test cases
	tests := []struct {
		description  string
		requestBody  User
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  User{"example@email.com", "fullname", 20},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid email",
			requestBody:  User{"invalid-email", "fullname", 20},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid fullname",
			requestBody:  User{"example@email.com", "12345", 20},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid age",
			requestBody:  User{"example@email.com", "fullname", -5},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
}
