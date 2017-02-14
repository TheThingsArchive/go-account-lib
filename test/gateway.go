// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import jwt "github.com/dgrijalva/jwt-go"

// GatewayClaims creates a jwt.Claims that represents the gateway
func GatewayClaims(id string, locationPublic, statusPublic bool) jwt.Claims {
	return jwt.MapClaims{
		"iss":             Issuer,
		"sub":             id,
		"type":            "gateway",
		"location_public": locationPublic,
		"status_public":   statusPublic,
	}
}

// GatewayToken creates a token that is singed by PrivateKey, and has the GatewayClaims
func GatewayToken(id string, locationPublic, statusPublic bool) string {
	return TokenFromClaims(GatewayClaims(id, locationPublic, statusPublic))
}
