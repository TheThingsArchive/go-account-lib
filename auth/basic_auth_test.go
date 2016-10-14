// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/assertions"
)

const (
	username = "foo"
	password = "secret"
)

func TestBasicAuthDecorate(t *testing.T) {
	a := New(t)

	strategy := BasicAuth(username, password)

	req, _ := http.NewRequest("GET", "/foo", nil)
	strategy.DecorateRequest(req)

	u, p, ok := req.BasicAuth()
	a.So(ok, ShouldBeTrue)
	a.So(u, ShouldEqual, username)
	a.So(p, ShouldEqual, password)
}

func TestBasicAuthWithScope(t *testing.T) {
	a := New(t)

	strategy := BasicAuth(username, password)

	withScope := strategy.WithScope("scope")
	a.So(withScope, ShouldEqual, strategy)
}
