// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
)

func TestMemoryStore(t *testing.T) {
	a := New(t)

	store := MemoryStore()

	// getting from an empty store should work
	res, err := store.Get(parent, scope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")

	// setting to a new store should work
	err = store.Set(parent, scopes, token, time.Second)
	a.So(err, ShouldBeNil)

	// getting from a not-so-new store should still work
	res, err = store.Get(parent, scope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, token)

	// trying to get an expired token should not work
	err = store.Set(parent, []string{"scope2"}, token, time.Duration(0))
	a.So(err, ShouldBeNil)

	res, err = store.Get(parent, otherScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")
}
