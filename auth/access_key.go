// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"
)

type accessKey struct {
	accessKey string
}

func (a *accessKey) DecorateRequest(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("key %s", a.accessKey))
}

// WithScope just returns itself
func (a *accessKey) WithScope(scope string) Strategy {
	return a
}

func AccessKey(s string) *accessKey {
	return &accessKey{
		accessKey: s,
	}
}
