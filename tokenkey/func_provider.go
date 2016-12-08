// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

import "fmt"

// TokenFunc returns a tokenKey and optionally renews it
type TokenFunc func(renew bool) (*TokenKey, error)

// funcProvider is a tokenkey Provider that executes a function which resturns a tokenkey
type funcProvider struct {
	funcs map[string]TokenFunc
}

// FuncProvider creates Provider that executes a function which resturns a tokenkey
func FuncProvider(funcs map[string]TokenFunc) Provider {
	return &funcProvider{
		funcs: funcs,
	}
}

func (p *funcProvider) Get(id string, renew bool) (*TokenKey, error) {
	fetch, ok := p.funcs[id]
	if !ok {
		return nil, fmt.Errorf("Auth server %s not registered", id)
	}
	return fetch(renew)
}

func (p *funcProvider) Update() error {
	for _, fetch := range p.funcs {
		_, err := fetch(true)
		if err != nil {
			return err
		}
	}
	return nil
}
