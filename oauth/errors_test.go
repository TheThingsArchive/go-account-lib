// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"errors"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestParseError(t *testing.T) {
	a := New(t)

	{
		err := fromError(errors.New(`oauth2: cannot fetch token: 400 Bad Request
Response: {"error":"invalid_grant","error_description":"Code not found"}`))

		a.So(err.Code, ShouldEqual, 400)
		a.So(err.Description, ShouldEqual, "Code not found")
	}

	{
		err := fromError(errors.New(`oauth2: cannot fetch token: 401 Unauthorized
Response: {"code":401,"error":"Invalid client id or secret"}`))
		a.So(err.Code, ShouldEqual, 401)
		a.So(err.Description, ShouldEqual, "Invalid client id or secret")
	}
}

func TestParseBadError(t *testing.T) {
	a := New(t)

	err := fromError(errors.New("foo"))

	a.So(err.Code, ShouldEqual, 500)
	a.So(err.Description, ShouldEqual, "foo")
}

func TestParseNil(t *testing.T) {
	a := New(t)

	err := fromError(nil)

	a.So(err, ShouldBeNil)
}
