// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"testing"

	"github.com/TheThingsNetwork/go-account-lib/tokens"
	. "github.com/smartystreets/assertions"
)

const (
	token      = "foo"
	scope      = "scope"
	otherToken = "bar"
)

func TestAccessTokenDecorate(t *testing.T) {
	a := New(t)

	strategy := AccessToken(token)
	req, _ := http.NewRequest("GET", "/foo", nil)

	strategy.DecorateRequest(req)
	a.So(req.Header.Get("Authorization"), ShouldEqual, "Bearer "+token)
}

func TestAccessTokenWithScope(t *testing.T) {
	a := New(t)

	strategy := AccessToken(token)

	withScope := strategy.WithScope(scope)
	a.So(withScope, ShouldNotEqual, strategy)

	at, _ := withScope.(*accessToken)

	a.So(at.scope, ShouldEqual, scope)
}

func TestAccessTokenWithManager(t *testing.T) {
	a := New(t)

	store := tokens.ConstStore(otherToken)

	strategy := AccessTokenWithManager(token, tokens.HTTPManager("server", token, store))
	withScope := strategy.WithScope(scope)

	req, _ := http.NewRequest("GET", "/foo", nil)

	withScope.DecorateRequest(req)
	a.So(req.Header.Get("Authorization"), ShouldEqual, "Bearer "+otherToken)
}
