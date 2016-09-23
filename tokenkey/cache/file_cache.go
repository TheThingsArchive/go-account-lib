// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// fileCache represents a cache that stores keys in memory
type fileCache struct {
	dirname string
}

// MemoryCache returns a cache that stores keys in memory
func FileCache(dirname string) *fileCache {
	return &fileCache{
		dirname: dirname,
	}
}

// filename builds the filename from the cache key
// and dirname
func (c *fileCache) filename(key string) string {
	return path.Join(c.dirname, fmt.Sprintf("auth-%s.pub", key))
}

// Get tries to read the filename
func (c *fileCache) Get(key string) ([]byte, error) {
	cached, err := ioutil.ReadFile(c.filename(key))
	if err != nil {
		// check the error
		if _, ok := err.(*os.PathError); !ok {
			return nil, err
		}
	}

	return cached, nil
}

// Set saves the data to the file determined by key
func (c *fileCache) Set(key string, data []byte) error {
	return ioutil.WriteFile(c.filename(key), data, 0644)
}
