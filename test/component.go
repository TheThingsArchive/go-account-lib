// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import jwt "github.com/dgrijalva/jwt-go"

// ComponentClaims creates a jwt.Claims that represents the component
func ComponentClaims(typ, id string) jwt.Claims {
	return jwt.MapClaims{
		"iss":  Issuer,
		"sub":  id,
		"type": typ,
	}
}

// ComponentToken creates a token that is singed by PrivateKey, and has the
// ComponentClaims
func ComponentToken(typ, id string) string {
	return TokenFromClaims(ComponentClaims(typ, id))
}
