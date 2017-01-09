// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

// AppRighter is the interface of everything that can hold rights to an app
type AppRighter interface {
	// AppRight checks wether or not the specified right is held on the app with the
	// specified appID
	AppRight(appID string, right Right) bool
}

// GatewayRighter is the interface of everything that can hold rights to a
// gateway
type GatewayRighter interface {
	// GatewayRight checks wether or not the specified right is held on the
	// gateway with the specified gatewayID
	GatewayRight(gatewayID string, right Right) bool
}

// ComponentRighter is the interface of everything that can hold rights to a
// component
type ComponentRighter interface {
	// ComponentRight checks wether or not the specified right is held on the
	// gateway with the specified gatewayID
	ComponentRight(gatewayID string, right Right) bool
}
