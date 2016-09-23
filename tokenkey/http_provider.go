// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/cache"
)

type httpProvider struct {
	servers map[string]string
	cache   cache.Cache
}

// HTTPProvider returns a new Provider that fetches the key from a HTTP resource
func HTTPProvider(servers map[string]string, cache cache.Cache) Provider {
	return &httpProvider{
		servers: servers,
		cache:   cache,
	}
}

func (p *httpProvider) Get(server string, renew bool) (*TokenKey, error) {
	data, _ := p.cache.Get(server)

	// Fetch token if there's a renew or if there's no key cached
	if renew || data == nil {
		fetched, err := p.fetch(server)
		if err != nil {
			// We don't have a key here
			return nil, err
		}

		data = fetched

		// We do not care about errors here
		_ = p.cache.Set(server, data)
	}

	var key TokenKey
	if err := json.Unmarshal(data, &key); err != nil {
		return nil, err
	}

	return &key, nil
}

func (p *httpProvider) Update() error {
	for server := range p.servers {
		data, err := p.fetch(server)
		if err != nil {
			return err
		}

		p.cache.Set(server, data)
	}
	return nil
}

func (p *httpProvider) fetch(server string) ([]byte, error) {
	url, ok := p.servers[server]
	if !ok {
		return nil, fmt.Errorf("Auth server %s not registered", server)
	}

	resp, err := http.Get(fmt.Sprintf("%s/key", url))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
