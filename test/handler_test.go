// handler_test.go
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	api "api-request-handler/internal/api"
	services "api-request-handler/internal/services"
)

func TestHandler(t *testing.T) {
	handler := api.NewHandler()

	req := httptest.NewRequest("GET", "/api/company/financials?companyId=test", nil)
	w := httptest.NewRecorder()

	handler.FinancialsHandler(services.DummyFinancialData)(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", w.Code)
	}
}
