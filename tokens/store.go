// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

// TokenStore is the interface of all the token storage engines
type TokenStore interface {

	// Get gets a token from the TokenStore, returning
	// the empty string if the token does not exist
	Get(sub string, kind string, ID string) (string, error)

	// Set saves a token, but only stores it for a given time
	Set(sub string, kind string, ID string, token string, TTL time.Duration) error
}
