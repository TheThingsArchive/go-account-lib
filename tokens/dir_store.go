// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/cache"
	"github.com/TheThingsNetwork/go-account-lib/claims"
)

type dirStore struct {
	cache cache.Cache
}

// DirStore creates a filestore that stores tokens in the
// specified directory
func DirStore(dirname string) TokenStore {
	return &dirStore{
		cache: cache.FileCacheWithNameFn(dirname, filename),
	}
}

// DirStoreWithNameFn creates a filestore that stores tokens in the
// specified directory under with a custom filename
func FileStoreWithNameFn(dirname string, nameFn func(string) string) TokenStore {
	return &dirStore{
		cache: cache.FileCacheWithNameFn(dirname, nameFn),
	}
}

// DirStoreWithFormat creates a filestore that stores tokens in the
// specified directory under with a custom filename
func DirStoreWithFormat(dirname, format string) TokenStore {
	return &dirStore{
		cache: cache.FileCacheWithFormat(dirname, format),
	}
}

// key creates a key for storing a token and scope by md5 hashing
// the pair
func (s *dirStore) key(parent, scope string) string {
	data := scope + "." + parent
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// filename is the default token filename
func filename(name string) string {
	return name + ".derived.token"
}

// Get gets the token and checks it's TTL
func (s *dirStore) Get(parent, scope string) (string, error) {
	key := s.key(parent, scope)
	data, err := s.cache.Get(key)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", nil
	}

	claims, err := claims.FromTokenWithoutValidation(string(data))
	if err != nil {
		return "", err
	}

	if claims.ExpiresAt == 0 {
		return "", nil
	}

	expires := time.Unix(claims.ExpiresAt, 0)
	if expires.Before(time.Now()) {
		return "", nil
	}

	return string(data), nil
}

// Set saves the token and sets its deadline
func (s *dirStore) Set(parent string, scopes []string, token string, TTL time.Duration) error {
	// store the token for every scope it has
	for _, scope := range scopes {
		key := s.key(parent, scope)
		s.cache.Set(key, []byte(token))
	}

	return nil
}
