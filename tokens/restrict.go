// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

type restrictRequest struct {
	Scope []string `json:"scope"`
}

type restrictResponse struct {
	AccessToken string `json:"access_token"`
}

// RestrictScope requests a new token with a different, more specific scope
func RestrictScope(server string, token string, scope []string) (string, error) {
	strategy := auth.AccessToken(token)
	req := restrictRequest{
		Scope: scope,
	}
	var res restrictResponse

	err := util.POST(server, strategy, "users/restrict-token", req, &res)
	if err != nil {
		return "", nil
	}

	return res.AccessToken, nil
}
