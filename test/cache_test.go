package test

import (
	"testing"
	"time"

	cache "api-request-handler/internal/cache"
)

func TestConcurrentCache(t *testing.T) {
	c := cache.NewConcurrentCache()

	key := "testKey"
	// Test the GetOrCompute method
	computed, err := c.GetOrCompute(key, func(k string) interface{} { return "computedValue" })
	if err != nil {
		t.Errorf("Error computing value: %v", err)
	}

	if computed.(string) != "computedValue" {
		t.Errorf("Expected 'computedValue', got %v", computed)
	}

	// Test expiry after cacheTTL
	time.Sleep(1 * time.Minute)
	// Call GetOrCompute again and expect the value to be recomputed.
	computed, err = c.GetOrCompute(key, func(k string) interface{} { return "newComputedValue" })
	if err != nil {
		t.Errorf("Error computing value: %v", err)
	}

	if computed.(string) != "newComputedValue" {
		t.Errorf("Expected 'newComputedValue', got %v", computed)
	}
}
