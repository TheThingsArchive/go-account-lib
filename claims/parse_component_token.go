// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"fmt"

	"github.com/TheThingsNetwork/go-account-lib/tokenkey"
)

const (
	RouterType  = "router"
	BrokerType  = "broker"
	HandlerType = "handler"
)

// FromComponentToken uses the tokenkey provider to parse and validate a token into its
// corresponding claims
func FromComponentToken(provider tokenkey.Provider, accessToken string) (*ComponentClaims, error) {
	claims := &ComponentClaims{}

	err := fromToken(provider, accessToken, claims)
	if err != nil {
		return nil, err
	}

	switch claims.Type {
	case RouterType, BrokerType, HandlerType:
		return claims, nil
	default:
		return nil, fmt.Errorf("Expected component token to have type 'router', 'broker' or 'handler', got '%s' instead", claims.Type)
	}
}

// FromTokenWithoutValidation parses a token into its corresponding claims,
// without checking the token signature
func FromComponentTokenWithoutValidation(accessToken string) (*ComponentClaims, error) {
	claims := &ComponentClaims{}
	return claims, fromTokenWithoutValidation(accessToken, claims)
}
