package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	baseURL := "http://localhost:8080"

	// Test 1: Public endpoint
	fmt.Println("=== Testing Public Endpoint ===")
	resp, err := http.Get(baseURL + "/api/v1/ping")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n\n", string(body))

	// Test 2: Create a test user (protected endpoint - will fail without auth)
	fmt.Println("=== Testing Protected Endpoint (should fail) ===")
	testUser := map[string]interface{}{
		"username": "testuser",
		"email":    "test@example.com",
		"clerk_id": "test_clerk_id_123",
		"bio":      "Test user bio",
	}

	jsonData, _ := json.Marshal(testUser)
	resp, err = http.Post(baseURL+"/api/v1/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n\n", string(body))

	// Test 3: List users (public endpoint)
	fmt.Println("=== Testing List Users ===")
	resp, err = http.Get(baseURL + "/api/v1/users")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n", string(body))
}
