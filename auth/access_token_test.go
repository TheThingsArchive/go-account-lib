// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"testing"
	"time"

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
	a.So(req.Header.Get("Authorization"), ShouldEqual, "bearer "+token)
}

func TestAccessTokenWithScope(t *testing.T) {
	a := New(t)

	strategy := AccessToken(token)

	withScope := strategy.WithScope(scope)
	a.So(withScope, ShouldNotEqual, strategy)

	at, _ := withScope.(*accessToken)

	a.So(at.scope, ShouldEqual, scope)
}

type constStore struct{}

func (s *constStore) Get(parent string, scope string) (string, error) {
	return otherToken, nil
}

func (s *constStore) Set(parent string, scope []string, token string, ttl time.Duration) error {
	return nil
}

func TestAccessTokenWithManager(t *testing.T) {
	a := New(t)

	store := &constStore{}

	strategy := AccessTokenWithManager(token, tokens.HTTPManager("server", token, store))
	withScope := strategy.WithScope(scope)

	req, _ := http.NewRequest("GET", "/foo", nil)

	withScope.DecorateRequest(req)
	a.So(req.Header.Get("Authorization"), ShouldEqual, "bearer "+otherToken)
}
