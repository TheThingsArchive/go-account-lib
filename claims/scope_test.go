// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"testing"

	"github.com/TheThingsNetwork/go-account-lib/scope"
	"github.com/smartystreets/assertions"
)

func TestClaimsAppsScope(t *testing.T) {
	a := assertions.New(t)

	a.So(none.HasScope(scope.Apps), assertions.ShouldBeFalse)
	a.So(empty.HasScope(scope.Apps), assertions.ShouldBeFalse)

	a.So(withAppAccess.HasScope(scope.Apps), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.HasScope(scope.Apps), assertions.ShouldBeFalse)
	a.So(withComponentAccess.HasScope(scope.Apps), assertions.ShouldBeFalse)

	a.So(withGatewaysScope.HasScope(scope.Apps), assertions.ShouldBeFalse)
	a.So(withComponentsScope.HasScope(scope.Apps), assertions.ShouldBeFalse)

	a.So(withAppsScope.HasScope(scope.Apps), assertions.ShouldBeTrue)
}
