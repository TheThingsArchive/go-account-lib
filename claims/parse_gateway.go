// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"fmt"

	"github.com/TheThingsNetwork/go-account-lib/tokenkey"
)

// FromGatewayToken uses the tokenkey provider to parse and validate a gateway token into its
// corresponding claims
func FromGatewayToken(provider tokenkey.Provider, gatewayToken string) (*GatewayClaims, error) {
	claims := &GatewayClaims{}
	if err := fromToken(provider, gatewayToken, claims); err != nil {
		return nil, err
	}

	if claims.Type != "gateway" {
		return nil, fmt.Errorf("Expected gateway token to have type 'gateway', but got '%s'", claims.Type)
	}

	return claims, nil
}

// FromGatewayTokenWithoutValidation parses a token into its corresponding
// gateway claims, without checking the token signature
func FromGatewayTokenWithoutValidation(gatewayToken string) (*GatewayClaims, error) {
	claims := &GatewayClaims{}
	return claims, fromTokenWithoutValidation(gatewayToken, claims)
}
