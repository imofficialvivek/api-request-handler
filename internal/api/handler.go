package api

import (
	"encoding/json"
	"net/http"

	cache "api-request-handler/internal/cache"
)

type Handler struct {
	cache *cache.ConcurrentCache
}

func NewHandler(c *cache.ConcurrentCache) *Handler {
	return &Handler{cache: c}
}

// FinancialsHandler handles the financial data API
func (h *Handler) FinancialsHandler(fetchData func(string) interface{}) http.HandlerFunc {
	return h.apiHandler("financials", fetchData)
}

// SalesDataHandler handles the sales data API
func (h *Handler) SalesDataHandler(fetchData func(string) interface{}) http.HandlerFunc {
	return h.apiHandler("sales", fetchData)
}

// EmployeeStatsHandler handles the employee stats API
func (h *Handler) EmployeeStatsHandler(fetchData func(string) interface{}) http.HandlerFunc {
	return h.apiHandler("employee", fetchData)
}

func (h *Handler) apiHandler(api string, fetchData func(string) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyID := r.URL.Query().Get("companyId")
		if companyID == "" {
			http.Error(w, "Missing companyId", http.StatusBadRequest)
			return
		}

		// Fetch or compute the data
		data, err := h.cache.GetOrCompute(companyID+"-"+api, fetchData)
		if err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}

		// Respond with the computed/cached data in JSON format
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			CompanyID string      `json:"companyId"`
			API       string      `json:"api"`
			Data      interface{} `json:"data"`
		}{
			CompanyID: companyID,
			API:       api,
			Data:      data,
		})
	}
}
