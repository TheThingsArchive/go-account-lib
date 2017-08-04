// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"testing"

	"github.com/smartystreets/assertions"
)

func TestClaimsAppAccess(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(empty.AppAccess(appID), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(withApplicationsScope.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(withComponentsScope.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(applicationRightsButNoScope.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.AppAccess(appID), assertions.ShouldBeFalse)

	// wrong scope
	a.So(withGatewayAccess.AppAccess(appID), assertions.ShouldBeFalse)
	a.So(withComponentAccess.AppAccess(appID), assertions.ShouldBeFalse)

	// correct scope
	a.So(withApplicationAccess.AppAccess(appID), assertions.ShouldBeTrue)
}

func TestClaimsGatewayAccess(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(empty.GatewayAccess(gatewayID), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(withApplicationsScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(withComponentsScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(applicationRightsButNoScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.GatewayAccess(gatewayID), assertions.ShouldBeFalse)

	// wrong scope
	a.So(withApplicationAccess.GatewayAccess(gatewayID), assertions.ShouldBeFalse)
	a.So(withComponentAccess.GatewayAccess(gatewayID), assertions.ShouldBeFalse)

	// correct scope
	a.So(withGatewayAccess.GatewayAccess(gatewayID), assertions.ShouldBeTrue)
}

func TestClaimsComponentAccess(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(empty.ComponentAccess(componentID), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(withApplicationsScope.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(withComponentsScope.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(applicationRightsButNoScope.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.ComponentAccess(componentID), assertions.ShouldBeFalse)

	// wrong scope
	a.So(withGatewayAccess.ComponentAccess(componentID), assertions.ShouldBeFalse)
	a.So(withApplicationAccess.ComponentAccess(componentID), assertions.ShouldBeFalse)

	// correct scope
	a.So(withComponentAccess.ComponentAccess(componentID), assertions.ShouldBeTrue)
}
