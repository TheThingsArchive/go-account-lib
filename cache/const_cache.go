// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

// constCache represents a cache that always returns the same data
type constCache struct {
	data []byte
}

// ConstCache creates a Cache that always returns the same data
func ConstCache(data []byte) Cache {
	return &constCache{
		data: data,
	}
}

// Get gets the same data every time
func (c *constCache) Get(key string) ([]byte, error) {
	return c.data, nil
}

// Set is a noop
func (c *constCache) Set(key string, data []byte) error {
	return nil
}
