// handler_test.go
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	api "api-request-handler/internal/api"
	cache "api-request-handler/internal/cache"
	services "api-request-handler/internal/services"
)

func TestHandler(t *testing.T) {
	c := cache.NewConcurrentCache()
	handler := api.NewHandler(c)

	req := httptest.NewRequest("GET", "/api/company/financials?companyId=test", nil)
	w := httptest.NewRecorder()

	handler.FinancialsHandler(services.DummyFinancialData)(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", w.Code)
	}
}
