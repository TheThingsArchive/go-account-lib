// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

// constProvider is a tokenkey Provider that always resturns the same tokenkey
type constProvider struct {
	key *TokenKey
}

// ConstProvider creates Provider that always resturns the same TokenKey
func ConstProvider(key, algorithm string) Provider {
	return &constProvider{
		key: &TokenKey{
			Algorithm: algorithm,
			Key:       key,
		},
	}
}

func (c *constProvider) Get(server string, renew bool) (*TokenKey, error) {
	return c.key, nil
}

func (c *constProvider) Update() error {
	return nil
}
