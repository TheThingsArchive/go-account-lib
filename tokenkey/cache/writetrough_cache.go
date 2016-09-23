// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

type writeTroughCache struct {
	memory *memoryCache
	file   *fileCache
}

// WriteTroughCache creates a cache that stores keys in memory
// and uses disk as fallback
func WriteTroughCache(dirname string) *writeTroughCache {
	return &writeTroughCache{
		memory: MemoryCache(),
		file:   FileCache(dirname),
	}
}

// Get gets the data from memory, and tries to read from disk
// if not there, caching the result
func (c *writeTroughCache) Get(key string) ([]byte, error) {
	// Try to read from memory cache
	data, err := c.memory.Get(key)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	// Try to read from cache file
	data, err = c.file.Get(key)
	if err != nil {
		return nil, err
	}

	// Store in memory cache and return
	err = c.memory.Set(key, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Set stores the data in memory and on disk
func (c *writeTroughCache) Set(key string, data []byte) error {
	err := c.memory.Set(key, data)
	if err != nil {
		return err
	}

	err = c.file.Set(key, data)
	if err != nil {
		return err
	}

	return nil
}
