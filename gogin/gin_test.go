package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {

	r := gin.Default()

	// Register the route
	r.POST("/create-user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
	})

	// Create a test user
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Convert user to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Could not marshal user: %v", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, "/create-user", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	resp := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(resp, req)

	// Check for expected status code
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.Code)
	}

	// Check response body
	expected := `{"message":"User created","user":{"name":"John Doe","age":30,"email":"john.doe@example.com"}}`
	if resp.Body.String() != expected {
		t.Errorf("Expected response body %s, got %s", expected, resp.Body.String())
	}
}
