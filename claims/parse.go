// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/TheThingsNetwork/go-account-lib/tokenkey"
	jwt "github.com/dgrijalva/jwt-go"
)

func FromToken(provider tokenkey.Provider, accessToken string) (claims *Claims, err error) {
	parsed, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if provider == nil {
			return nil, errors.New("No token provider configured")
		}

		k, err := provider.Get(claims.Issuer, false)
		if err != nil {
			return nil, err
		}

		if k.Algorithm != token.Header["alg"] {
			return nil, fmt.Errorf("expected algorithm %v but got %v", k.Algorithm, token.Header["alg"])
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(k.Key))
		if err != nil {
			return nil, err
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to parse token: %s", err)
	}

	if !parsed.Valid {
		return nil, fmt.Errorf("token not valid or expired")
	}

	return claims, nil
}

func FromTokenWithoutValidation(accessToken string) (*Claims, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("Invalid access token segments")
	}

	segment, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return nil, err
	}

	var claims Claims
	err = json.Unmarshal(segment, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
