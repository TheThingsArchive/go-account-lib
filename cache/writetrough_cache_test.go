// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestWriteTroughCache(t *testing.T) {
	a := New(t)

	dir, err := ioutil.TempDir("", "test")
	a.So(err, ShouldBeNil)
	defer os.RemoveAll(dir)

	cache := WriteTroughCache(dir)

	got, err := cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(got, ShouldBeNil)

	err = cache.Set(key, data)
	a.So(err, ShouldBeNil)

	got, err = cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)

	// A second cache should pick up on the changes
	otherCache := WriteTroughCache(dir)
	got, err = otherCache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)

	err = otherCache.Set(key, data)
	a.So(err, ShouldBeNil)

	got, err = otherCache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)

	// Should still be in memory even if filesystem
	// does not have it
	os.RemoveAll(dir)

	got, err = cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)
}
