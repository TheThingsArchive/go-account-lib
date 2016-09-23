// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

// emptyCache represents a cache that never stores anything
type emptyCache struct{}

// EmptyCache represents a cache that never stores anything
var EmptyCache = &emptyCache{}

// Get gets no data every time
func (c *emptyCache) Get(key string) ([]byte, error) {
	return nil, nil
}

// Set is a noop
func (c *emptyCache) Set(key string, data []byte) error {
	return nil
}
