// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import "net/http"

// Strategy represents an authorization method
type Strategy interface {
	// DecorateRequest decorates an HTTP
	// request with an authorization method
	// for the given scope
	DecorateRequest(request *http.Request)

	// WithScope transforms the strategy into
	// a new Strategy that uses the requested scope
	WithScope(scope string) Strategy
}
