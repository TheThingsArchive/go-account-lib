// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cache

import (
	"testing"

	. "github.com/smartystreets/assertions"
)

var (
	key  = "key"
	data = []byte{0x01, 0x02, 0x03}
)

func TestEmptyCache(t *testing.T) {
	a := New(t)

	got, err := EmptyCache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(got, ShouldBeNil)

	err = EmptyCache.Set(key, data)
	a.So(err, ShouldBeNil)

	got, err = EmptyCache.Get(key)
	a.So(err, ShouldBeNil)
	a.So(got, ShouldBeNil)
}
