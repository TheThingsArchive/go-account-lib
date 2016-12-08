// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

import (
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestFuncProvider(t *testing.T) {
	a := New(t)

	tokenKey := new(TokenKey)
	provider := FuncProvider(map[string]TokenFunc{
		"id": func(bool) (*TokenKey, error) {
			return tokenKey, nil
		},
	})

	{
		_, err := provider.Get("bad-id", false)
		a.So(err, ShouldNotBeNil)
	}

	{
		res, err := provider.Get("id", false)
		a.So(err, ShouldBeNil)
		a.So(res, ShouldEqual, tokenKey)
	}

	{
		err := provider.Update()
		a.So(err, ShouldBeNil)
	}

}
