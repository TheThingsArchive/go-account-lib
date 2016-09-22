// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import "time"

type tokenWithDeadline struct {
	Token    string
	Deadline time.Time
}

type memoryStore struct {
	tokens map[string]tokenWithDeadline
}

// key builds the key that will be used in the tokens map
func (s *memoryStore) key(server, sub, scope string) string {
	return server + ":" + sub + ":" + scope
}

// Get gets the token and checks it's TTL
func (s *memoryStore) Get(server, sub, scope string) (string, error) {
	key := s.key(server, sub, scope)
	tok, ok := s.tokens[key]

	// token not set
	if !ok {
		return "", nil
	}

	// token set but expired, remove it from the map
	if time.Now().After(tok.Deadline) {
		delete(s.tokens, key)
		return "", nil
	}

	return tok.Token, nil
}

// Set saves the token and sets its deadline
func (s *memoryStore) Set(server, sub string, scopes []string, token string, TTL time.Duration) error {
	deadline := time.Now().Add(TTL)
	tok := tokenWithDeadline{
		Token:    token,
		Deadline: deadline,
	}

	// store the token for every scope it has
	for _, scope := range scopes {
		key := s.key(server, sub, scope)
		s.tokens[key] = tok
	}

	return nil
}

func MemoryStore() memoryStore {
	return memoryStore{
		tokens: map[string]tokenWithDeadline{},
	}
}
