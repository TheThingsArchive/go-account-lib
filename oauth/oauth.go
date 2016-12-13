// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/cache"
	"github.com/TheThingsNetwork/go-account-lib/util"
	"github.com/TheThingsNetwork/go-utils/log"
	"golang.org/x/oauth2"
)

type Config struct {
	Server string
	Client *Client
	cache  cache.Cache
}

// Client represents a client for the OAuth 2.0 flow
type Client struct {
	ID           string
	Secret       string
	RedirectURL  string
	ExtraHeaders map[string]string
}

// OAuth creates a new 3-legged OAuth client
func OAuth(server string, client *Client) *Config {
	return &Config{
		Server: server,
		Client: client,
		cache:  cache.EmptyCache,
	}
}

// OAuth creates a new 3-legged OAuth client that uses a cache to cache token
// exchanges
func OAuthWithCache(server string, client *Client, cache cache.Cache) *Config {
	return &Config{
		Server: server,
		Client: client,
		cache:  cache,
	}
}

// o.getConfig builds the oauth2 config for an OAuth client
func (o *Config) getConfig() *oauth2.Config {
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
func (o *Config) getKeyConfig() *oauth2.Config {
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
func (c *Config) getContext() context.Context {
	client := &http.Client{
		Transport: util.NewRoundTripper(log.Get(), c.Client.ExtraHeaders),
	}

	return context.WithValue(context.Background(), oauth2.HTTPClient, client)
}

// Exchange exchanges an OAuth 2.0 Authorization Code for an oauth2.Token
func (o *Config) Exchange(code string) (*oauth2.Token, error) {
	config := o.getConfig()
	token, err := config.Exchange(o.getContext(), code)
	return token, fromError(err)
}

// PasswordCredentialsToken gets an oauth2.Token from username and password
func (o *Config) PasswordCredentialsToken(username, password string) (*oauth2.Token, error) {
	config := o.getConfig()
	token, err := config.PasswordCredentialsToken(o.getContext(), username, password)
	return token, fromError(err)
}

// TokenSource creates oauth2.TokenSource from an oauht2.Token
func (o *Config) TokenSource(token *oauth2.Token) oauth2.TokenSource {
	config := o.getConfig()
	return config.TokenSource(o.getContext(), token)
}

// ExchangeAppKeyForToken exchanges an application Access Key for an equivalent
func (o *Config) ExchangeAppKeyForToken(appID, accessKey string) (*oauth2.Token, error) {
	// application Access Token
	config := o.getKeyConfig()

	token, err := getTokenFromCache(o.cache, appID, accessKey)
	if err != nil {
		return nil, err
	}

	if token != nil {
		return token, nil
	}

	token, err = config.PasswordCredentialsToken(o.getContext(), appID, accessKey)
	if err != nil {
		return nil, fromError(err)
	}

	// ignore errors when saving to cache
	_ = saveTokenToCache(o.cache, appID, accessKey, token)

	return token, nil
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page that asks for permissions for the required scopes explicitly.
func (o *Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	config := o.getConfig()
	return config.AuthCodeURL(state, opts...)
}
