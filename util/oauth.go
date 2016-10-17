// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package util

import (
	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// MakeConfig creates an oauth.Config object based on the necessary parameters,
// redirectURL can be left empty if not needed
func MakeConfig(server string, clientID string, clientSecret string, redirectURL string) oauth2.Config {
	endpoint := oauth2.Endpoint{
		TokenURL: fmt.Sprintf("%s/users/token", server),
		AuthURL:  fmt.Sprintf("%s/users/authorize", server),
	}

	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		RedirectURL:  redirectURL,
	}
}

// MakeKeyConfig creates an oauth2.Config that can be used to exchange an app
// access key for an access token
func MakeKeyConfig(server string, clientID string, clientSecret string) oauth2.Config {
	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/api/v2/applications/token", server),
		},
	}
}

// ExchangeKeyForToken exchanges and app access key for an app access token
func ExchangeKeyForToken(config oauth2.Config, appID, accessKey string) (*oauth2.Token, error) {
	return config.PasswordCredentialsToken(context.TODO(), appID, accessKey)
}
