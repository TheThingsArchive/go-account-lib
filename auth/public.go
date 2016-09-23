// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import "net/http"

// public is the auth strategy that does not add any authorization
type public struct{}

// DecorateRequest is a noop
func (p *public) DecorateRequest(req *http.Request) {}

// WithScope just returns itself
func (p *public) WithScope(scope string) Strategy {
	return p
}

// Public is the auth strategy that does not add any authorization
var Public = &public{}
