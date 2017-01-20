// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"testing"

	"github.com/TheThingsNetwork/go-account-lib/oauth"
	"github.com/TheThingsNetwork/go-account-lib/util"
	. "github.com/smartystreets/assertions"
)

func TestStatusCode(t *testing.T) {
	a := New(t)

	{
		code := StatusCode(&util.HTTPError{
			Code: 403,
		})
		a.So(code, ShouldEqual, 403)
	}

	{
		code := StatusCode(util.HTTPError{
			Code: 401,
		})

		a.So(code, ShouldEqual, 401)
	}

	{
		code := StatusCode(&oauth.Error{
			Code: 400,
		})

		a.So(code, ShouldEqual, 400)
	}

	{
		code := StatusCode(fmt.Errorf("ok"))
		a.So(code, ShouldEqual, 500)
	}
}
