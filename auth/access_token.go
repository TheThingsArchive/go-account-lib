// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/tokens"
)

// accessToken is the Strategy that uses access tokens
// to authorize
type accessToken struct {
	accessToken string
	scope       string
	manager     tokens.Manager
}

// DecorateRequest gets the correct access token and uses that to
// decorate the request
func (a *accessToken) DecorateRequest(req *http.Request) {
	var token string

	// try to get new token and fall back to parent token
	if a.scope != "" {
		token, _ = a.manager.TokenForScope(a.scope)
	}

	if token == "" {
		token = a.accessToken
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

// WithScope creates a new Strategy with
func (a *accessToken) WithScope(scope string) Strategy {
	return &accessToken{
		accessToken: a.accessToken,
		manager:     a.manager,
		scope:       scope,
	}
}

// AccessToken returns an authorization strategy that uses the specified
// access token to authorize
func AccessToken(s string) Strategy {
	return &accessToken{
		accessToken: s,
		manager:     tokens.HTTPManager("", s, tokens.NullStore),
	}
}

// AccessTokenWithManager returns an authorization strategy that uses the specified
// access token to authorize and that uses the specified manager to
// fetch new tokens if required.
func AccessTokenWithManager(s string, manager tokens.Manager) Strategy {
	return &accessToken{
		accessToken: s,
		manager:     manager,
	}
}
