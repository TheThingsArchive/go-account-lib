// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import (
	"github.com/TheThingsNetwork/go-account-lib/scope"
	"github.com/TheThingsNetwork/ttn/core/types"
	jwt "github.com/dgrijalva/jwt-go"
)

// UserClaims creates user claims for top level scopes (like apps, profile, gateways)
func UserClaims(id string, scopes []string) jwt.Claims {
	return jwt.MapClaims{
		"iss":             Issuer,
		"sub":             id,
		"type":            "user",
		"interchangeable": true,
		"scope":           scopes,
	}
}

// UserToken creates user claims for top level scopes (like apps, profile, gateways)
func UserToken(id string, scopes []string) string {
	return TokenFromClaims(UserClaims(id, scopes))
}

// DerivedUserClaims creates a user token with derived claims (like app:foo)
func DerivedUserClaims(id string, apps map[string][]types.Right, gateways map[string][]types.Right, components map[string][]types.Right) jwt.Claims {
	scopes := make([]string, 0, len(apps)+len(gateways)+len(components))

	for id := range apps {
		scopes = append(scopes, scope.Application(id))
	}

	for id := range gateways {
		scopes = append(scopes, scope.Gateway(id))
	}

	for id := range components {
		scopes = append(scopes, scope.Component(id))
	}

	return jwt.MapClaims{
		"iss":             Issuer,
		"sub":             id,
		"type":            "user",
		"interchangeable": false,
		"scope":           scopes,
		"apps":            apps,
		"gateways":        gateways,
		"components":      components,
	}
}

// DerivedUserToken creates a user token with derived claims (like app:foo)
func DerivedUserToken(id string, apps map[string][]types.Right, gateways map[string][]types.Right, components map[string][]types.Right) string {
	return TokenFromClaims(DerivedUserClaims(id, apps, gateways, components))
}
