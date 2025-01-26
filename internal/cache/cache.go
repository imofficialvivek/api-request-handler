package cache

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentCache struct {
	data       map[string]entry
	mutex      sync.RWMutex
	locks      sync.Map
	waitingMap sync.Map
}

type entry struct {
	value  interface{}
	expiry time.Time
}

const cacheTTL = 1 * time.Minute

func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		data: make(map[string]entry),
	}
}

func (c *ConcurrentCache) GetOrCompute(key string, compute func(string) interface{}) (interface{}, error) {
	// Try to load existing data (skip if already cached)
	c.mutex.RLock()
	if e, found := c.data[key]; found && time.Now().Before(e.expiry) {
		c.mutex.RUnlock()
		return e.value, nil
	}
	c.mutex.RUnlock()

	// Lock mutex per key for synchronized computation
	mutexInterface, _ := c.locks.LoadOrStore(key, &sync.Mutex{})
	mutex := mutexInterface.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	// Check again if the computation already ran or cached before waiting
	c.mutex.RLock()
	if e, found := c.data[key]; found && time.Now().Before(e.expiry) {
		c.mutex.RUnlock()
		return e.value, nil
	}
	c.mutex.RUnlock()

	// Print that calculation is being initiated
	fmt.Printf("Calculating initial data for: %s\n", key)

	// Create or get the existing channel for waiting requests
	waitChannel := make(chan struct{})
	existingChannelInterface, _ := c.waitingMap.LoadOrStore(key, waitChannel)
	if existingChannelInterface != waitChannel {
		// If there is an existing channel, wait for it to close after computation
		<-existingChannelInterface.(chan struct{})
	}

	// Launch the computation in a goroutine
	var result interface{}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		result = compute(key)
	}()

	// Wait for goroutine to complete
	wg.Wait()

	// After computation completes, handle cache update and signaling
	c.mutex.Lock()
	c.data[key] = entry{value: result, expiry: time.Now().Add(cacheTTL)}
	c.mutex.Unlock()

	// Notify all waiting requests
	close(waitChannel)

	// Remove the key from waitingMap after the work is done
	c.waitingMap.Delete(key)

	return result, nil
}
