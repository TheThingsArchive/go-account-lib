// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
)

func TestNullStore(t *testing.T) {
	a := New(t)
	store := NullStore{}

	// getting from a ne  sotre should work
	res, err := store.Get("sub", "kind", "ID")
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")

	// setting to a new store should work
	err = store.Set("sub", "kind", "ID", "token", time.Second)
	a.So(err, ShouldBeNil)

	// getting from a not-so-new store should still work
	res, err = store.Get("sub", "kind", "ID")
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")
}
