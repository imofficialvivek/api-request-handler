package main

import (
	"fmt"
	"net/http"

	api "api-request-handler/internal/api"
	services "api-request-handler/internal/services"
)

func main() {
	// Initialize API handler
	handler := api.NewHandler()

	// Define route for financial data
	http.HandleFunc("/api/financials", handler.FinancialsHandler(services.DummyFinancialData))

	// Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
