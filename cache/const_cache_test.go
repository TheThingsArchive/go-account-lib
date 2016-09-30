// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestConstCache(t *testing.T) {
	a := New(t)

	cache := ConstCache(data)

	// Getting shouyld return the preloaded data.
	got, err := cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)

	// Setting should give no error.
	err = cache.Set(key, []byte{0x03, 0x04})
	a.So(err, ShouldBeNil)

	// Getting after setting should return the preloaded data.
	got, err = cache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(bytes.Equal(data, got), ShouldBeTrue)
}
