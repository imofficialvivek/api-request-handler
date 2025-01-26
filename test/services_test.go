package test

import (
	"testing"

	services "api-request-handler/internal/services"
)

func TestInitialData(t *testing.T) {
	data := services.InitialData("id")
	if data == nil {
		t.Error("Initial data returned nil")
	}

	// Additional checks for specific fields
	expectedCompany := "Kelp"
	expectedLocation := "Andheri"

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		t.Error("Initial data did not return a map[string]interface{}")
		return
	}

	if company, found := dataMap["company"]; !found || company != expectedCompany {
		t.Errorf("Expected company: %v, got: %v", expectedCompany, company)
	}

	if location, found := dataMap["location"]; !found || location != expectedLocation {
		t.Errorf("Expected location: %v, got: %v", expectedLocation, location)
	}
}
