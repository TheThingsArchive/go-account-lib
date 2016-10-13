// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import jwt "github.com/dgrijalva/jwt-go"

// TokenFromClaims creates a JWT with the specified claims, signed by the privateKey
func TokenFromClaims(claims jwt.Claims) string {
	builder := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(PrivateKey)
	if err != nil {
		panic(err)
	}

	token, err := builder.SignedString(key)
	if err != nil {
		panic(err)
	}

	return token
}
