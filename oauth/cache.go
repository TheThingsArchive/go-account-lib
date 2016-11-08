// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/cache"

	"golang.org/x/oauth2"
)

func key(appID, accessKey string) string {
	return appID + ":" + accessKey
}

func getTokenFromCache(cache cache.Cache, appID, accessKey string) (*oauth2.Token, error) {
	data, err := cache.Get(key(appID, accessKey))
	if err != nil {
		return nil, err
	}

	if data != nil {
		var token oauth2.Token

		err = json.Unmarshal(data, &token)
		if err != nil {
			return nil, err
		}

		fmt.Println(token.Expiry, time.Now())

		// only return token if not expired
		if token.Expiry.After(time.Now()) {
			return &token, nil
		}
	}

	return nil, nil
}

func saveTokenToCache(cache cache.Cache, appID, accessKey string, token *oauth2.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	return cache.Set(key(appID, accessKey), data)
}
