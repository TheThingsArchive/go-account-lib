// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package keys

import "strings"

// KeyIssuer returns the account server issuer id that issued the key,
// if no issuer id is found, it returns the empty string.
func KeyIssuer(accessKey string) string {
	parts := strings.Split(accessKey, ".")
	if len(parts) != 2 {
		return ""
	}

	return parts[0]
}
