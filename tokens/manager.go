// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"time"

	"github.com/TheThingsNetwork/go-account-lib/claims"
)

type Manager struct {
	server string
	token  string
	store  TokenStore
}

// ManagerWithStore creates a new token manager that uses the specified store
// to store and retrieve tokens
func ManagerWithStore(server, token string, store TokenStore) *Manager {
	return &Manager{
		server: server,
		token:  token,
		store:  store,
	}
}

// NullManager creates a new token manager that uses the NullStore
// to store and retrieve tokens
func NullManager(server, token string) *Manager {
	return ManagerWithStore(server, token, &NullStore)
}

// MemoryManager creates a new token manager that uses the MemoryStore
// to store and retrieve tokens
func MemoryManager(server, token string) *Manager {
	return ManagerWithStore(server, token, MemoryStore())
}

// TokenForScopes returns a token that will work for the specified scopes
func (m *Manager) TokenForScopes(scopes []string) (string, error) {
	// try to get existing token
	token, err := m.store.Get(m.token, scope)
	if err != nil {
		return "", err
	}

	// return token if it exists
	if token != "" {
		return token, nil
	}

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

// TokenForScope returns a token that will work for the specified scope
func (m *Manager) TokenForScope(scope string) (string, error) {
	return m.TokenForScopes([]string{scope})
}

// TokenForApp returns a token that works for the specified app
func (m *Manager) TokenForApp(appID string) (string, error) {
	return m.TokenForScope(claims.AppScope + ":" + appID)
}

// TokenForGateway returns a token that works for the specified gateway
func (m *Manager) TokenForGateway(gatewayID string) (string, error) {
	return m.TokenForScope(claims.GatewayScope + ":" + gatewayID)
}

// TokenForComponent returns a token that works for the specified gateway
func (m *Manager) TokenForComponent(componentID string) (string, error) {
	return m.TokenForScope(claims.ComponentScope + ":" + componentID)
}
