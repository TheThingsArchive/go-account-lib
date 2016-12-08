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

// fromToken parser a token given the tokenkey provider into the desired claims
// structure
func fromToken(provider tokenkey.Provider, token string, claims ClaimsWithIssuer) error {
	parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (key interface{}, err error) {
		if provider == nil {
			return nil, errors.New("No token provider configured")
		}

		k, err := provider.Get(claims.GetIssuer(), false)
		if err != nil {
			return nil, err
		}

		if k.Algorithm != token.Header["alg"] {
			return nil, fmt.Errorf("expected algorithm %v but got %v", k.Algorithm, token.Header["alg"])
		}

		switch k.Algorithm {
		case "RS256":
			key, err = jwt.ParseRSAPublicKeyFromPEM([]byte(k.Key))
		case "ES256":
			key, err = jwt.ParseECPublicKeyFromPEM([]byte(k.Key))
		default:
			err = fmt.Errorf("Token provider returned public key with unknown algorithm %s", k.Algorithm)
		}
		if err != nil {
			return nil, err
		}

		return key, nil
	})

	if err != nil {
		return fmt.Errorf("unable to parse token: %s", err)
	}

	if !parsed.Valid {
		return fmt.Errorf("token not valid or expired")
	}

	return nil
}

// FromTokenWithoutValidation parses a token into the desired claims structure,
// without checking the token signature
func fromTokenWithoutValidation(token string, claims jwt.Claims) error {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("Invalid access token segments")
	}

	segment, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return err
	}

	err = json.Unmarshal(segment, claims)
	if err != nil {
		return err
	}

	return nil
}
