// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

// TokenKey is the data returned by the token key provider
type TokenKey struct {
	Algorithm string `json:"algorithm"`
	Key       string `json:"key"`
}

// Provider represents a provider of the token key
type Provider interface {
	Get(server string, renew bool) (*TokenKey, error)
	Update() error
}
