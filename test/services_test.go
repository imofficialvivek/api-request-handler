package test

import (
	"testing"

	services "api-request-handler/internal/services"
)

func TestDummyServices(t *testing.T) {
	if data := services.DummyFinancialData("id"); data == nil {
		t.Error("DummyFinancialData returned nil")
	}
}
