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
	nameFn  func(string) string
}

const defaultFormat = "ttn-%s.data"

// FileCache returns a cache that stores keys on filesystem
func FileCache(dirname string) Cache {
	return FileCacheWithFormat(dirname, defaultFormat)
}

// FileCacheWithNameFn creates a FileCache that has a custom function to generate
// filenames
func FileCacheWithNameFn(dirname string, filename func(string) string) Cache {
	return &fileCache{
		dirname: dirname,
		nameFn:  filename,
	}
}

// FileCacheWithFormat creates a FileCache that uses a format string to generate
// filenames
func FileCacheWithFormat(dirname string, format string) Cache {
	return FileCacheWithNameFn(dirname, func(key string) string {
		return fmt.Sprintf(format, key)
	})
}

// filename creates the full filename for key based on configuration
func (c *fileCache) filename(key string) string {
	return path.Join(c.dirname, c.nameFn(key))
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
	err := os.MkdirAll(c.dirname, 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.filename(key), data, 0644)
}
