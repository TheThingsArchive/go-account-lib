// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import "net/http"

type basicAuth struct {
	username string
	password string
}

//  DecorateRequest decorates a request with the access key
func (a *basicAuth) DecorateRequest(req *http.Request) {
	req.SetBasicAuth(a.username, a.password)
}

// WithScope just returns itself
func (a *basicAuth) WithScope(scope string) Strategy {
	return a
}

// AccessKey creates a authorization strategy that uses
// the specified access key
func BasicAuth(username, password string) Strategy {
	return &basicAuth{
		username: username,
		password: password,
	}
}
