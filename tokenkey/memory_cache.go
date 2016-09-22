// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

// memoryCache represents a cache that stores keys in memory
type memoryCache struct {
	cache map[string][]byte
}

// MemoryCache returns a cache that stores keys in memory
func MemoryCache() *memoryCache {
	return &memoryCache{
		cache: map[string][]byte{},
	}
}

func (c *memoryCache) Get(key string) ([]byte, error) {
	cached, ok := c.cache[key]
	if ok {
		return cached, nil
	}

	return nil, nil
}

func (c *memoryCache) Set(key string, data []byte) error {
	c.cache[key] = data
	return nil
}
