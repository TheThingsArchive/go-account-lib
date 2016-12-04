// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"reflect"
	"testing"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/test"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/assertions"
)

func buildComponentClaims(TTL time.Duration, typ string) *ComponentClaims {
	return &ComponentClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "foo",
			ExpiresAt: time.Now().Add(TTL).Unix(),
		},
		Type: typ,
	}
}

func TestParseValidComponenToken(t *testing.T) {
	a := New(t)

	claims := buildComponentClaims(time.Hour, "router")
	token := test.TokenFromClaims(claims)

	parsed, err := FromComponentToken(test.Provider, token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}

func TestParseComponentExpired(t *testing.T) {
	a := New(t)

	claims := buildComponentClaims(-1*time.Hour, "handler")
	token := test.TokenFromClaims(claims)

	_, err := FromComponentToken(test.Provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseComponentInvalidType(t *testing.T) {
	a := New(t)

	claims := buildComponentClaims(time.Hour, "invalid")
	token := test.TokenFromClaims(claims)

	_, err := FromComponentToken(test.Provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseComponentWithoutValidation(t *testing.T) {
	a := New(t)

	{
		claims := buildComponentClaims(time.Hour, "broker")
		token := test.TokenFromClaims(claims)

		parsed, err := FromComponentTokenWithoutValidation(token)
		a.So(err, ShouldBeNil)
		a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
	}

	{
		claims := buildComponentClaims(time.Hour, "bla")
		token := test.TokenFromClaims(claims)

		parsed, err := FromComponentTokenWithoutValidation(token)
		a.So(err, ShouldBeNil)
		a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
	}
}
