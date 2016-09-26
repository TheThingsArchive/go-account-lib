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

//  DecorateRequest decorates a request with the access key
func (a *accessKey) DecorateRequest(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("key %s", a.accessKey))
}

// WithScope just returns itself
func (a *accessKey) WithScope(scope string) Strategy {
	return a
}

// AccessKey creates a authorization strategy that uses
// the specified access key
func AccessKey(s string) Strategy {
	return &accessKey{
		accessKey: s,
	}
}
