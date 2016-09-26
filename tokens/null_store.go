// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

// The NullStore will not save tokens, but will never give an
// error and so is perfect to use in situations where storing tokens
// is not necessary or impossible.
type nullStore struct{}

// Get always returns no token
func (s *nullStore) Get(parent, scope string) (string, error) {
	return "", nil
}

// Set accepts a token but does nothing with it
func (s *nullStore) Set(parent string, scope []string, token string, TTL time.Duration) error {
	return nil
}

// NullStore is a TokenStore that does not store tokens at all
var NullStore = &nullStore{}
