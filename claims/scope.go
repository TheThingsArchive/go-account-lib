// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

// HasScope checks if the token has the specified scope
func (claims *Claims) HasScope(scope string) bool {
	return claims != nil && contains(claims.Scope, scope)
}

// hasScopedID checks if the token has the specified composite scope
// For example `claims.hasScopedID("apps", "foo")` checks if the scope `"apps:foo"` is
// present on `claims`.
func (claims *Claims) hasScopedID(scope, ID string) bool {
	return claims.HasScope(scope + ":" + ID)
}
