// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import jwt "github.com/dgrijalva/jwt-go"

type location struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

// GatewayClaims creates a jwt.Claims that represents the gateway
func GatewayClaims(id, frequencyPlan string, lat, lng float64, locationPublic, statusPublic bool) jwt.Claims {
	return jwt.MapClaims{
		"iss":             Issuer,
		"sub":             id,
		"frequency_plan":  frequencyPlan,
		"location_public": locationPublic,
		"status_public":   statusPublic,
		"location": location{
			Longitude: lng,
			Latitude:  lat,
		},
	}
}

// GatewayToken creates a token that is singed by PrivateKey, and has the GatewayClaims
func GatewayToken(id, frequencyPlan string, lat, lng float64, locationPublic, statusPublic bool) string {
	return TokenFromClaims(GatewayClaims(id, frequencyPlan, lat, lng, locationPublic, statusPublic))
}
