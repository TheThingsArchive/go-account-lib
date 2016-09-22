// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

const appID = "foo"
const gatewayID = "foo"
const componentID = "foo"
const right = "settings"
const otherRight = "delete"

var none *Claims
var empty = Claims{}

var withAppAccess = Claims{
	Scope: []string{"apps:" + appID},
	Apps: map[string][]string{
		appID: []string{right},
	},
}

var withAppsScope = Claims{
	Scope: []string{"apps"},
}

var appRightsButNoScope = Claims{
	Apps: map[string][]string{
		appID: []string{right},
	},
}

var withGatewayAccess = Claims{
	Scope: []string{"gateways:" + gatewayID},
	Gateways: map[string][]string{
		gatewayID: []string{right},
	},
}

var withGatewaysScope = Claims{
	Scope: []string{"gateways"},
}

var gatewayRightsButNoScope = Claims{
	Gateways: map[string][]string{
		gatewayID: []string{right},
	},
}

var withComponentAccess = Claims{
	Scope: []string{"components:" + componentID},
	Components: map[string][]string{
		componentID: []string{right},
	},
}

var withComponentsScope = Claims{
	Scope: []string{"components"},
}

var componentRightsButNoScope = Claims{
	Components: map[string][]string{
		componentID: []string{right},
	},
}
