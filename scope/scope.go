// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package scope

const (
	// Apps is the scope name for apps
	Apps = "apps"

	// Gateways is the scope name for gateways
	Gateways = "gateways"

	// Components is the scope name for components
	Components = "components"

	// Profile is the scope name for profiles
	Profile = "profile"
)

// App returns the scope for the app with the specific ID
func App(ID string) string {
	return Apps + ":" + ID
}

// Gateway returns the scope for the gateway with the specific ID
func Gateway(ID string) string {
	return Gateways + ":" + ID
}

// Component returns the scope for the component with the specific ID
func Component(ID string) string {
	return Components + ":" + ID
}
