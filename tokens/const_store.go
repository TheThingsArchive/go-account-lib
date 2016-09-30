// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

// constStore represents a store that always returns the same token
type constStore struct {
	token string
}

// ConstStore creates a TokensStore that always returns the same token
func ConstStore(token string) TokenStore {
	return &constStore{
		token: token,
	}
}

// Get always returns the initial token
func (s *constStore) Get(parent, scope string) (string, error) {
	return s.token, nil
}

// Set accepts a token but does nothing with it
func (s *constStore) Set(parent string, scope []string, token string, TTL time.Duration) error {
	return nil
}
