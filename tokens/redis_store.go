// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"crypto/md5"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/claims"

	redis "gopkg.in/redis.v5"
)

const prefix = "token"
const TTLMargin = 5 * time.Second

// hash hashes the parent token to make shorter redis keys
func hash(parent string) string {
	res := md5.Sum([]byte(parent))
	return string(res[:])
}

// key creates a key for a given parent token and scope
func key(parent, scope string) string {
	return prefix + ":" + scope + ":" + hash(parent)
}

// redisStore is a TokenStore that stores the tokens in redis
type redisStore struct {
	client *redis.Client
}

// RedisStore returns a TokenStore that stores tokens in redis
func RedisStore(client *redis.Client) TokenStore {
	if client == nil {
		panic("RedisStore: redis client is nil")
	}
	return &redisStore{
		client: client,
	}
}

// Get gets a token that was derived from parent for the specified scope
// if it exists in redis
func (s *redisStore) Get(parent, scope string) (string, error) {
	token, err := s.client.Get(key(parent, scope)).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return token, err
}

// Set saves the token that is derived from parent and that has the specified
// scopes to the redis store
func (s *redisStore) Set(parent string, scopes []string, token string, TTL time.Duration) (err error) {
	// use expiry from parent token
	claims, err := claims.FromTokenWithoutValidation(parent)
	if err != nil {
		return err
	}

	// get ttl from token
	ttl := time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) - TTLMargin

	// if TTL is explicitly set, use that one
	if TTL != time.Duration(0) {
		ttl = TTL
	}

	if ttl < time.Duration(0) {
		return nil
	}

	for _, scope := range scopes {
		err = s.client.Set(key(parent, scope), token, TTL).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
