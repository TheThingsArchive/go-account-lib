// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import "github.com/TheThingsNetwork/ttn/core/types"

const appID = "foo"
const gatewayID = "foo"
const componentID = "foo"
const right = types.Right("settings")
const otherRight = "delete"

var none *Claims
var empty = Claims{}

var withApplicationAccess = Claims{
	Scope: []string{"apps:" + appID},
	Apps: map[string][]types.Right{
		appID: []types.Right{right},
	},
}

var withApplicationsScope = Claims{
	Scope: []string{"apps"},
}

var applicationRightsButNoScope = Claims{
	Apps: map[string][]types.Right{
		appID: []types.Right{right},
	},
}

var withGatewayAccess = Claims{
	Scope: []string{"gateways:" + gatewayID},
	Gateways: map[string][]types.Right{
		gatewayID: []types.Right{right},
	},
}

var withGatewaysScope = Claims{
	Scope: []string{"gateways"},
}

var gatewayRightsButNoScope = Claims{
	Gateways: map[string][]types.Right{
		gatewayID: []types.Right{right},
	},
}

var withComponentAccess = Claims{
	Scope: []string{"components:" + componentID},
	Components: map[string][]types.Right{
		componentID: []types.Right{right},
	},
}

var withComponentsScope = Claims{
	Scope: []string{"components"},
}

var componentRightsButNoScope = Claims{
	Components: map[string][]types.Right{
		componentID: []types.Right{right},
	},
}
