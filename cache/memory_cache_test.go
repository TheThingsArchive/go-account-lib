// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestMemoryCache(t *testing.T) {
	a := New(t)

	cache := MemoryCache()

	got, err := cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(got, ShouldBeNil)

	err = cache.Set(key, data)
	a.So(err, ShouldBeNil)

	got, err = cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)
}
