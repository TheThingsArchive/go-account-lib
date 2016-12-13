// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"os"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/tokens"
	"github.com/TheThingsNetwork/go-utils/handlers/cli"
	"github.com/TheThingsNetwork/go-utils/log"
	wrap "github.com/TheThingsNetwork/go-utils/log/apex"
	apex "github.com/apex/log"
)

// Account is a client to an account server
type Account struct {
	server  string
	auth    auth.Strategy
	ctx     log.Interface
	headers map[string]string
}

// New creates a new account that will use no authentication
func New(server string) *Account {
	return &Account{
		server: server,
		auth:   auth.Public,
		ctx: wrap.Wrap(&apex.Logger{
			Handler: cli.New(os.Stdout),
		}),
		headers: map[string]string{},
	}
}

// WithLogger sets the logger that the account will use to log warnings
func (a *Account) WithLogger(logger log.Interface) *Account {
	a.ctx = logger
	return a
}

// WithAuth sets the authentication strategy the account will use
func (a *Account) WithAuth(strategy auth.Strategy) *Account {
	a.auth = strategy
	return a
}

// WithHeader adds the header to every request to the account server
func (a *Account) WithHeader(name, value string) *Account {
	a.headers[name] = value
	return a
}

// NewWithAccessToken creates a new account client that will use the
// accessToken to make requests to the specified account server
func NewWithAccessToken(server, accessToken string) *Account {
	return New(server).WithAuth(auth.AccessToken(accessToken))
}

// NewWithManager creates a new account client that will use the
// accessToken to make requests to the specified account server
// and the manager to request new tokens for different scopes
func NewWithManager(server, accessToken string, manager tokens.Manager) *Account {
	return New(server).WithAuth(auth.AccessTokenWithManager(accessToken, manager))
}

// NewWithKey creates an account client that uses an accessKey to
// authenticate
func NewWithKey(server, accessKey string) *Account {
	return New(server).WithAuth(auth.AccessKey(accessKey))
}

// NewWithBasicAuth creates an account client that uses basic authentication
func NewWithBasicAuth(server, username, password string) *Account {
	return New(server).WithAuth(auth.BasicAuth(username, password))
}

// NewWithPublic creates an account client that does not use authentication
func NewWithPublic(server string) *Account {
	return New(server).WithAuth(auth.Public)
}
