// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestPublicDecorate(t *testing.T) {
	a := New(t)

	req, _ := http.NewRequest("GET", "/foo", nil)
	Public.DecorateRequest(req)

	a.So(req.Header.Get("Authorization"), ShouldEqual, "")
}

func TestPublicWithScope(t *testing.T) {
	a := New(t)

	withScope := Public.WithScope("scope")
	a.So(withScope, ShouldEqual, Public)
}
