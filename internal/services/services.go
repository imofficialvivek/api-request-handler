package services

func DummyFinancialData(companyID string) interface{} {
	return map[string]interface{}{
		"revenue": 200000,
		"profit":  100000,
	}
}

func DummySalesData(companyID string) interface{} {
	return map[string]interface{}{
		"units_sold": 1000,
		"returns":    200,
	}
}
