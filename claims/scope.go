// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

// HasScope checks if the token has the specified scope
func (claims *Claims) HasScope(scope string) bool {
	return claims != nil && contains(claims.Scope, scope)
}
