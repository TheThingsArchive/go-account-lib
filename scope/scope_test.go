// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package scope

import (
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestApplication(t *testing.T) {
	a := New(t)
	a.So(Application("foo"), ShouldEqual, "apps:foo")
}

func TestGateway(t *testing.T) {
	a := New(t)
	a.So(Gateway("foo"), ShouldEqual, "gateways:foo")
}

func TestComponent(t *testing.T) {
	a := New(t)
	a.So(Component("foo"), ShouldEqual, "components:foo")
}
