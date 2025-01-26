package main

import (
	"fmt"
	"net/http"

	api "api-request-handler/internal/api"
	cache "api-request-handler/internal/cache"
	services "api-request-handler/internal/services"
)

func main() {
	fmt.Println("------Starting API Request Handler------")

	// Initialize cache
	c := cache.NewConcurrentCache()

	// Initialize API handlers
	handler := api.NewHandler(c)

	// Define routes
	http.HandleFunc("/api/company/financials", handler.FinancialsHandler(services.DummyFinancialData))
	http.HandleFunc("/api/sales/data", handler.SalesDataHandler(services.DummySalesData))
	http.HandleFunc("/api/employee/stats", handler.EmployeeStatsHandler(services.DummyEmployeeStats))

	// Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
