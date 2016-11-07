// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package keys

import (
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestKeyIssuer(t *testing.T) {

	a := New(t)

	// key without issuer
	{
		issuer := KeyIssuer("foo")
		a.So(issuer, ShouldBeEmpty)
	}

	// key with issuer
	{
		issuer := KeyIssuer("foo.bar")
		a.So(issuer, ShouldEqual, "foo")
	}

}
