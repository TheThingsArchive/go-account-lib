// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"testing"

	"github.com/smartystreets/assertions"
)

func TestClaimsAppRight(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(empty.AppRight(appID, right), assertions.ShouldBeFalse)

	// has only global scope
	a.So(withApplicationScope.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(withComponentsScope.AppRight(appID, right), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(applicationRightsButNoScope.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.AppRight(appID, right), assertions.ShouldBeFalse)

	// wrong right
	a.So(withApplicationAccess.AppRight(appID, otherRight), assertions.ShouldBeFalse)

	// wrong scope/rights
	a.So(withGatewayAccess.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.AppRight(appID, otherRight), assertions.ShouldBeFalse)
	a.So(withComponentAccess.AppRight(appID, right), assertions.ShouldBeFalse)
	a.So(withComponentAccess.AppRight(appID, otherRight), assertions.ShouldBeFalse)

	// correct scope and rights
	a.So(withApplicationAccess.AppRight(appID, right), assertions.ShouldBeTrue)
}

func TestClaimsGatewayRight(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(empty.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)

	// has only global scope
	a.So(withApplicationScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(withComponentsScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(applicationRightsButNoScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)

	// wrong right
	a.So(withApplicationAccess.GatewayRight(gatewayID, otherRight), assertions.ShouldBeFalse)

	// wrong scope/rights
	a.So(withApplicationAccess.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)
	a.So(withApplicationAccess.GatewayRight(gatewayID, otherRight), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.GatewayRight(gatewayID, otherRight), assertions.ShouldBeFalse)
	a.So(withComponentAccess.GatewayRight(gatewayID, otherRight), assertions.ShouldBeFalse)
	a.So(withComponentAccess.GatewayRight(gatewayID, right), assertions.ShouldBeFalse)

	// correct scope and rights
	a.So(withGatewayAccess.GatewayRight(gatewayID, right), assertions.ShouldBeTrue)
}

func TestClaimsComponentRight(t *testing.T) {
	a := assertions.New(t)

	// empty claims
	a.So(none.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(empty.ComponentRight(componentID, right), assertions.ShouldBeFalse)

	// has only global scope
	a.So(withApplicationScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(withGatewaysScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(withComponentsScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)

	// no app:id scope present
	a.So(applicationRightsButNoScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(gatewayRightsButNoScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(componentRightsButNoScope.ComponentRight(componentID, right), assertions.ShouldBeFalse)

	// wrong right
	a.So(withApplicationAccess.ComponentRight(componentID, otherRight), assertions.ShouldBeFalse)

	// wrong scope/rights
	a.So(withApplicationAccess.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(withApplicationAccess.ComponentRight(componentID, otherRight), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.ComponentRight(componentID, right), assertions.ShouldBeFalse)
	a.So(withGatewayAccess.ComponentRight(componentID, otherRight), assertions.ShouldBeFalse)
	a.So(withComponentAccess.ComponentRight(componentID, otherRight), assertions.ShouldBeFalse)

	// correct scope and rights
	a.So(withComponentAccess.ComponentRight(componentID, right), assertions.ShouldBeTrue)
}
