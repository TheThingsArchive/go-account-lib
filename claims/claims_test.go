// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import "github.com/TheThingsNetwork/go-account-lib/rights"

const appID = "foo"
const gatewayID = "foo"
const componentID = "foo"
const right = rights.Right("settings")
const otherRight = "delete"

var none *Claims
var empty = Claims{}

var withAppAccess = Claims{
	Scope: []string{"apps:" + appID},
	Apps: map[string][]rights.Right{
		appID: []rights.Right{right},
	},
}

var withAppsScope = Claims{
	Scope: []string{"apps"},
}

var appRightsButNoScope = Claims{
	Apps: map[string][]rights.Right{
		appID: []rights.Right{right},
	},
}

var withGatewayAccess = Claims{
	Scope: []string{"gateways:" + gatewayID},
	Gateways: map[string][]rights.Right{
		gatewayID: []rights.Right{right},
	},
}

var withGatewaysScope = Claims{
	Scope: []string{"gateways"},
}

var gatewayRightsButNoScope = Claims{
	Gateways: map[string][]rights.Right{
		gatewayID: []rights.Right{right},
	},
}

var withComponentAccess = Claims{
	Scope: []string{"components:" + componentID},
	Components: map[string][]rights.Right{
		componentID: []rights.Right{right},
	},
}

var withComponentsScope = Claims{
	Scope: []string{"components"},
}

var componentRightsButNoScope = Claims{
	Components: map[string][]rights.Right{
		componentID: []rights.Right{right},
	},
}
