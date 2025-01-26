package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	api "api-request-handler/internal/api"
	cache "api-request-handler/internal/cache"
	services "api-request-handler/internal/services"
)

func TestHandler(t *testing.T) {
	// Initialize the cache and handler
	c := cache.NewConcurrentCache()
	handler := api.NewHandler(c)

	// Test the FinancialsHandler with Initial data
	req := httptest.NewRequest("GET", "/api/company/financials?companyId=test", nil)
	w := httptest.NewRecorder()

	handler.FinancialsHandler(services.InitialData)(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", w.Code)
	}

	// Decode the actual response body
	var actualBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	// Construct expected body
	expectedBody := map[string]interface{}{
		"companyId": "test",
		"api":       "financials",
		"data": map[string]interface{}{
			"company":  "Kelp",
			"location": "Andheri",
		},
	}

	// Comparing actual vs. expected
	if actualBody["companyId"] != expectedBody["companyId"] ||
		actualBody["api"] != expectedBody["api"] {
		t.Errorf("Mismatch in metadata. Expected %v, got %v", expectedBody, actualBody)
	}

	data, ok := actualBody["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Data field is not a map[string]interface{}")
	}

	expectedData := expectedBody["data"].(map[string]interface{})
	if data["company"] != expectedData["company"] || data["location"] != expectedData["location"] {
		t.Errorf("Mismatch in data. Expected %v, got %v", expectedData, data)
	}
}
