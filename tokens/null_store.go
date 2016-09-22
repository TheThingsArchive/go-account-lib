// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

// The NullStore will not save tokens, but will never give an
// error and so is perfect to use in situations where storing tokens
// is not necessary or impossible.
type NullStore struct{}

// Get always returns no token
func (s *NullStore) Get(sub, kind, ID string) (string, error) {
	return "", nil
}

// Set accepts a token but does nothing with it
func (s *NullStore) Set(sub, kind, ID, token string, TTL time.Duration) error {
	return nil
}
