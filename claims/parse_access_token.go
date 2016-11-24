// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import "github.com/TheThingsNetwork/go-account-lib/tokenkey"

// FromToken uses the tokenkey provider to parse and validate a token into its
// corresponding claims
func FromToken(provider tokenkey.Provider, accessToken string) (*Claims, error) {
	claims := &Claims{}
	return claims, fromToken(provider, accessToken, claims)
}

// FromTokenWithoutValidation parses a token into its corresponding claims,
// without checking the token signature
func FromTokenWithoutValidation(accessToken string) (*Claims, error) {
	claims := &Claims{}
	return claims, fromTokenWithoutValidation(accessToken, claims)
}
