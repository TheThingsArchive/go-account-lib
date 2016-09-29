// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type fileTokenStore struct {
	filename string
}

func FileStore(filename string) TokenStore {
	return &fileTokenStore{
		filename: filename,
	}
}

type fileToken struct {
	Parent  string    `json:"parent_token"`
	Scope   string    `json:"scope"`
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (s *fileTokenStore) readTokens() ([]fileToken, error) {
	// read the token file
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []fileToken{}, nil
		}
		return nil, err
	}

	now := time.Now()
	res := make([]fileToken, 0)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// unmarshal the token
		var token fileToken
		err := json.Unmarshal([]byte(line), &token)
		if err != nil {
			return nil, err
		}

		// only append tokens that are not expired
		if token.Expires.After(now) {
			res = append(res, token)
		}
	}

	return res, nil
}

func (s *fileTokenStore) writeTokens(tokens []fileToken) error {
	now := time.Now()
	lines := make([]string, 0)
	for _, token := range tokens {
		if token.Token != "" && token.Expires.After(now) {
			data, err := json.Marshal(token)
			if err != nil {
				return err
			}
			lines = append(lines, string(data))
		}
	}
	return ioutil.WriteFile(s.filename, []byte(strings.Join(lines, "\n")), 0600)
}

func (s *fileTokenStore) Get(parent string, scope string) (string, error) {
	tokens, err := s.readTokens()
	if err != nil {
		return "", err
	}

	for _, token := range tokens {
		if token.Parent == parent && token.Scope == scope {
			return token.Token, nil
		}
	}
	return "", nil
}

func (s *fileTokenStore) Set(parent string, scopes []string, tok string, TTL time.Duration) error {
	tokens, err := s.readTokens()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, scope := range scopes {
		updated := false
		newToken := fileToken{
			Parent:  parent,
			Token:   tok,
			Scope:   scope,
			Expires: now.Add(TTL),
		}
		for i, token := range tokens {
			if token.Parent == parent && token.Scope == scope {
				tokens[i] = newToken
				updated = true
			}
		}

		if !updated {
			tokens = append(tokens, newToken)
		}
	}

	return s.writeTokens(tokens)
}
