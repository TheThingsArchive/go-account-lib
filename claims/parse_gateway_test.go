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

func buildGatewayClaims(TTL time.Duration, typ string) *GatewayClaims {
	return &GatewayClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "foo",
			ExpiresAt: time.Now().Add(TTL).Unix(),
		},
		Type:           typ,
		LocationPublic: true,
		StatusPublic:   false,
	}
}

func TestParseValidGatewayToken(t *testing.T) {
	a := New(t)

	claims := buildGatewayClaims(time.Hour, "gateway")
	token := test.TokenFromClaims(claims)

	parsed, err := FromGatewayToken(test.Provider, token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}

func TestParseGatewayExpired(t *testing.T) {
	a := New(t)

	claims := buildGatewayClaims(-1*time.Hour, "gateway")
	token := test.TokenFromClaims(claims)

	_, err := FromGatewayToken(test.Provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseGatewayInvalidType(t *testing.T) {
	a := New(t)

	claims := buildGatewayClaims(time.Hour, "invalid")
	token := test.TokenFromClaims(claims)

	_, err := FromGatewayToken(test.Provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseGatewayWithoutValidation(t *testing.T) {
	a := New(t)

	{
		claims := buildGatewayClaims(time.Hour, "gateway")
		token := test.TokenFromClaims(claims)

		parsed, err := FromGatewayTokenWithoutValidation(token)
		a.So(err, ShouldBeNil)
		a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
	}

	{
		claims := buildGatewayClaims(time.Hour, "bla")
		token := test.TokenFromClaims(claims)

		parsed, err := FromGatewayTokenWithoutValidation(token)
		a.So(err, ShouldBeNil)
		a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
	}
}
