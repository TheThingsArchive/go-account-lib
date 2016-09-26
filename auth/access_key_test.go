// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/assertions"
)

const (
	key = "foo"
)

func TestAccessKeyDecorate(t *testing.T) {
	a := New(t)

	strategy := AccessKey(key)

	req, _ := http.NewRequest("GET", "/foo", nil)
	strategy.DecorateRequest(req)

	a.So(req.Header.Get("Authorization"), ShouldEqual, "key "+key)
}

func TestAccessKeyWithScope(t *testing.T) {
	a := New(t)

	strategy := AccessKey(key)

	withScope := strategy.WithScope("scope")
	a.So(withScope, ShouldEqual, strategy)
}
