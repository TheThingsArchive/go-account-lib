// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"fmt"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/util"
	"github.com/TheThingsNetwork/go-utils/log"
	"golang.org/x/oauth2"
)

// ExchangeGatewayKeyForToken exchanges an application Access Key for an equivalent
func (o *Config) ExchangeGatewayKeyForToken(gatewayID, gatewayKey string) (*oauth2.Token, error) {
	strategy := auth.AccessKey(gatewayKey)
	token := &oauth2.Token{}
	err := util.GET(log.Get(), o.Server, strategy, fmt.Sprintf("/api/v2/gateways/%s/token", gatewayID), token)
	return token, err
}
