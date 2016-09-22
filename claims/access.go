// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

// AppAccess checks if the token grants access to the specified app
func (claims *Claims) AppAccess(appID string) bool {
	return claims.hasScopedID(AppScope, appID)
}

// GatewayAccess checks if the token grants access to the specified Gateway
func (claims *Claims) GatewayAccess(gatewayID string) bool {
	return claims.hasScopedID(GatewayScope, gatewayID)
}

// ComponentAccess checks if the token grants access to the specified Component
func (claims *Claims) ComponentAccess(componentID string) bool {
	return claims.hasScopedID(ComponentScope, componentID)
}
