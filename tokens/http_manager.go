// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

type httpManager struct {
	server string
	token  string
	store  TokenStore
}

// HTTPManager creates a new token manager that uses the specified store
// to store and retrieve tokens
func HTTPManager(server, token string, store TokenStore) Manager {
	return &httpManager{
		server: server,
		token:  token,
		store:  store,
	}
}

// TokenForScopes returns a token that will work for the specified scopes
func (m *httpManager) TokenForScope(scope string) (string, error) {
	// try to get existing token
	token, err := m.store.Get(m.token, scope)
	if err != nil {
		return "", err
	}

	// return token if it exists
	if token != "" {
		return token, nil
	}

	scopes := []string{scope}

	// token did not exist, get one from the server
	restricted, err := RestrictScope(m.server, m.token, scopes)
	if err != nil {
		return "", err
	}

	// store the new token
	err = m.store.Set(m.token, scopes, restricted, time.Hour)
	if err != nil {
		return "", err
	}

	return restricted, nil
}
