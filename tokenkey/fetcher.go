// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

// Fetcher represents something that can fetch a TokenKey
type Fetcher interface {
	Fetch(string) (*TokenKey, error)
}
