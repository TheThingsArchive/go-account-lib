// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import jwt "github.com/dgrijalva/jwt-go"

const AppScope = "apps"
const GatewayScope = "gateways"
const ComponentScope = "components"

type Claims struct {
	jwt.StandardClaims
	Client     string              `json:"client"`
	Scope      []string            `json:"scope"`
	Apps       map[string][]string `json:"apps,omitempty"`
	Gateways   map[string][]string `json:"gateways,omitempty"`
	Components map[string][]string `json:"components,omitempty"`
}
