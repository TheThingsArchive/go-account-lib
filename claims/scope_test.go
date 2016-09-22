// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"testing"

	"github.com/smartystreets/assertions"
)

func TestClaimsAppsScope(t *testing.T) {
	a := assertions.New(t)

	a.So(none.HasScope(AppScope), assertions.ShouldBeFalse)
	a.So(empty.HasScope(AppScope), assertions.ShouldBeFalse)

	a.So(withAppAccess.HasScope(AppScope), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.HasScope(AppScope), assertions.ShouldBeFalse)
	a.So(withComponentAccess.HasScope(AppScope), assertions.ShouldBeFalse)

	a.So(withGatewaysScope.HasScope(AppScope), assertions.ShouldBeFalse)
	a.So(withComponentsScope.HasScope(AppScope), assertions.ShouldBeFalse)

	a.So(withAppsScope.HasScope(AppScope), assertions.ShouldBeTrue)
}
