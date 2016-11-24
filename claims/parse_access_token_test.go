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

func buildClaims(TTL time.Duration) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "foo",
			ExpiresAt: time.Now().Add(TTL).Unix(),
		},
		Email:    "example@example.com",
		Client:   "foo",
		Scope:    []string{"apps"},
		Type:     "handler",
		Username: "example",
		Name: struct {
			First string `json:"first"`
			Last  string `json:"last"`
		}{
			First: "John",
			Last:  "Doe",
		},
	}
}

func TestParseValidToken(t *testing.T) {
	a := New(t)

	claims := buildClaims(time.Hour)
	token := test.TokenFromClaims(claims)

	parsed, err := FromToken(test.Provider, token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}

func TestParseExpired(t *testing.T) {
	a := New(t)

	claims := buildClaims(-1 * time.Hour)
	token := test.TokenFromClaims(claims)

	_, err := FromToken(test.Provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseWithoutValidation(t *testing.T) {
	a := New(t)

	claims := buildClaims(time.Hour)
	token := test.TokenFromClaims(claims)

	parsed, err := FromTokenWithoutValidation(token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}
