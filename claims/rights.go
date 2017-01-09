// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import "github.com/TheThingsNetwork/go-account-lib/rights"

// AppRight checks if the token grants the specified right for
// the app with the specified ID
func (claims *Claims) AppRight(appID string, right rights.Right) bool {
	return claims.AppAccess(appID) && containsRight(claims.Apps[appID], right)
}

// GatewayRight checks if the token grants the specified right for
// the Gateway with the specified ID
func (claims *Claims) GatewayRight(gatewayID string, right rights.Right) bool {
	return claims.GatewayAccess(gatewayID) && containsRight(claims.Gateways[gatewayID], right)
}

// ComponentRight checks if the token grants the specified right for
// the Component with the specified ID
func (claims *Claims) ComponentRight(componentID string, right rights.Right) bool {
	return claims.ComponentAccess(componentID) && containsRight(claims.Components[componentID], right)
}
