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
func (s *memoryStore) key(sub, kind, ID string) string {
	return sub + ":" + kind + ":" + ID
}

// Get gets the token and checks it's TTL
func (s *memoryStore) Get(sub, kind, ID string) (string, error) {
	key := s.key(sub, kind, ID)
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
func (s *memoryStore) Set(sub, kind, ID, token string, TTL time.Duration) error {
	key := s.key(sub, kind, ID)
	deadline := time.Now().Add(TTL)

	s.tokens[key] = tokenWithDeadline{
		Token:    token,
		Deadline: deadline,
	}

	return nil
}

func MemoryStore() memoryStore {
	return memoryStore{
		tokens: map[string]tokenWithDeadline{},
	}
}
