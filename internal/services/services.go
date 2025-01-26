package services

func DummyFinancialData(companyID string) interface{} {
	return map[string]interface{}{
		"revenue": 200000,
		"profit":  100000,
	}
}
