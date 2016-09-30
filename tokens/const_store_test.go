// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
)

const (
	otherToken = "othertoken"
)

func TestConstStore(t *testing.T) {
	a := New(t)
	store := ConstStore(token)

	// getting from a ne  sotre should work
	res, err := store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, token)

	// setting to a new store should work
	err = store.Set(parent, scopes, otherToken, time.Second)
	a.So(err, ShouldBeNil)

	// getting from a not-so-new store should still work
	res, err = store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, token)
}
