// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

// Manager represent all types that can manage tokens
type Manager interface {
	// Get a token for the specified scope
	TokenForScope(scope string) (string, error)
}
