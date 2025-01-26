package cache

type ConcurrentCache struct {
	data map[string]interface{}
}

func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		data: make(map[string]interface{}),
	}
}
