// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"fmt"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/util"
	"github.com/TheThingsNetwork/go-utils/log"
	"golang.org/x/oauth2"
)

type tok struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint64 `json:"expires_in"`
}

func (tok *tok) Token() *oauth2.Token {
	if tok == nil {
		return nil
	}
	return &oauth2.Token{
		AccessToken: tok.AccessToken,
		Expiry:      time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second),
	}
}

// ExchangeGatewayKeyForToken exchanges an application Access Key for an equivalent
func (o *Config) ExchangeGatewayKeyForToken(gatewayID, gatewayKey string) (*oauth2.Token, error) {
	strategy := auth.AccessKey(gatewayKey)
	token := &tok{}
	err := util.GET(log.Get(), o.Server, strategy, fmt.Sprintf("/api/v2/gateways/%s/token", gatewayID), o.Client.ExtraHeaders, token)
	return token.Token(), err
}
