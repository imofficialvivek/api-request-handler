package api

import (
	"fmt"
	"net/http"
)

type Handler struct{}

// Create a new handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// FinancialsHandler handles the financial data API
func (h *Handler) FinancialsHandler(fetchData func(string) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyID := r.URL.Query().Get("companyId")
		if companyID == "" {
			http.Error(w, "Missing companyId", http.StatusBadRequest)
			return
		}

		// Fetch financial data
		data := fetchData(companyID)

		// Return the financial data as response
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%v", data)
	}
}
