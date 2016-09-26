// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

// memoryCache represents a cache that stores keys in memory
type memoryCache struct {
	cache map[string][]byte
}

// MemoryCache returns a cache that stores keys in memory
func MemoryCache() Cache {
	return &memoryCache{
		cache: map[string][]byte{},
	}
}

// Get gets the data from memory if it exists
func (c *memoryCache) Get(key string) ([]byte, error) {
	cached, ok := c.cache[key]
	if ok {
		return cached, nil
	}

	return nil, nil
}

// Set saves the data to memory
func (c *memoryCache) Set(key string, data []byte) error {
	c.cache[key] = data
	return nil
}
