// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

// Cache represents a cache
type Cache interface {
	// Get gets the data, returning nil if it isn't cached
	Get(key string) ([]byte, error)

	// Set saves data,
	Set(key string, data []byte) error
}
