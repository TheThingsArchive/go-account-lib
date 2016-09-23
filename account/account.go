// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/tokens"
)

// Account is a client to an account server
type Account struct {
	server string
	auth   auth.Strategy
}

// New creates a new account client that will use the
// accessToken to make requests to the specified account server
func New(server, accessToken string) *Account {
	return &Account{
		server: server,
		auth:   auth.AccessToken(accessToken),
	}
}

// NewWithManager creates a new account client that will use the
// accessToken to make requests to the specified account server
// and the manager to request new tokens for different scopes
func NewWithManager(server, accessToken string, manager *tokens.Manager) *Account {
	return &Account{
		server: server,
		auth:   auth.AccessTokenWithManager(accessToken, manager),
	}
}

// NewWithKey creates an account client that uses an accessKey to
// authenticate
func NewWithKey(server, accessKey string) *Account {
	return &Account{
		server: server,
		auth:   auth.AccessKey(accessKey),
	}
}

// NewWithPublic creates an account client that does not use authentication
func NewWithPublic(server string) *Account {
	return &Account{
		server: server,
		auth:   auth.Public,
	}
}
