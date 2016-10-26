// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type oauth struct {
	Server string
	Client *Client
}

// Client represents a client for the OAuth 2.0 flow
type Client struct {
	ID          string
	Secret      string
	RedirectURL string
}

// OAuth creates a new 3-legged OAuth client
func OAuth(server string, client *Client) *oauth {
	return &oauth{
		Server: server,
		Client: client,
	}
}

// o.getConfig builds the oauth2 config for an OAuth client
func (o *oauth) getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     o.Client.ID,
		ClientSecret: o.Client.Secret,
		RedirectURL:  o.Client.RedirectURL,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/users/token", o.Server),
			AuthURL:  fmt.Sprintf("%s/users/authorize", o.Server),
		},
	}
}

// getKeyConfig builds the oauth2 config for an OAuth client to exchange and app
// key for an app token
func (o *oauth) getKeyConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     o.Client.ID,
		ClientSecret: o.Client.Secret,
		RedirectURL:  o.Client.RedirectURL,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/api/v2/applications/token", o.Server),
		},
	}
}

// getContext builds a new context for use in an oauth exchange
func getContext() context.Context {
	return context.Background()
}

// Exchange exchanges an OAuth 2.0 Authorization Code for an oauth2.Token
func (o *oauth) Exchange(code string) (*oauth2.Token, error) {
	config := o.getConfig()
	return config.Exchange(getContext(), code)
}

// PasswordCredentialsToken gets an oauth2.Token from username and password
func (o *oauth) PasswordCredentialsToken(username, password string) (*oauth2.Token, error) {
	config := o.getConfig()
	return config.PasswordCredentialsToken(getContext(), username, password)
}

// TokenSource creates oauth2.TokenSource from an oauht2.Token
func (o *oauth) TokenSource(token *oauth2.Token) oauth2.TokenSource {
	config := o.getConfig()
	return config.TokenSource(getContext(), token)
}

// ExchangeAppKeyForToken exchanges an application Access Key for an equivalent
func (o *oauth) ExchangeAppKeyForToken(appID, accessKey string) (*oauth2.Token, error) {
	// application Access Token
	config := o.getKeyConfig()
	return config.PasswordCredentialsToken(getContext(), appID, accessKey)
}
