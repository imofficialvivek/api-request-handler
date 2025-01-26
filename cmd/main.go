package main

import (
	"fmt"
	"net/http"

	api "api-request-handler/internal/api"
	cache "api-request-handler/internal/cache"
	services "api-request-handler/internal/services"
)

func main() {
	// Initialize cache
	c := cache.NewConcurrentCache()

	// Initialize API handlers
	handler := api.NewHandler(c)

	// Define route for financial data
	http.HandleFunc("/api/company/financials", handler.FinancialsHandler(services.DummyFinancialData))

	// Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
