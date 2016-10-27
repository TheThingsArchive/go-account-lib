// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"testing"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/test"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/assertions"
)

func makeToken(scope []string, dur time.Duration) string {
	claims := jwt.MapClaims(map[string]interface{}{
		"sub":   "tester",
		"exp":   int((time.Now().Add(dur)).Unix()) * 1000,
		"scope": scope,
	})

	return test.TokenFromClaims(claims)
}

var (
	testScope   = "apps:foo"
	otherScope  = "gateways:bar"
	scopes      = []string{testScope}
	otherScopes = []string{otherScope}
	bothScopes  = []string{testScope, otherScope}

	parent      = makeToken(scopes, time.Second)
	otherParent = makeToken(otherScopes, time.Second)

	token      = makeToken(scopes, time.Second)
	otherToken = makeToken(otherScopes, time.Second)
)

func TestNullStore(t *testing.T) {
	a := New(t)
	store := NullStore

	// getting from a ne  sotre should work
	res, err := store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")

	// setting to a new store should work
	err = store.Set(parent, scopes, token, time.Second)
	a.So(err, ShouldBeNil)

	// getting from a not-so-new store should still work
	res, err = store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")
}
