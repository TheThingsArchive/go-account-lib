// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"
)

type accessToken struct {
	accessToken string
}

func (a *accessToken) DecorateRequest(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", a.accessToken))
}

func AccessToken(s string) *accessToken {
	return &accessToken{
		accessToken: s,
	}
}
