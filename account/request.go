// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"io"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

func (a *Account) get(strategy auth.Strategy, URI string, res interface{}) error {
	return util.GET(a.ctx, a.server, strategy, URI, a.headers, res)
}

func (a *Account) gets(strategy auth.Strategy, URI string) (io.ReadCloser, error) {
	return util.GETBody(a.ctx, a.server, strategy, URI, a.headers)
}

func (a *Account) put(strategy auth.Strategy, URI string, body, res interface{}) error {
	return util.PUT(a.ctx, a.server, strategy, URI, a.headers, body, res)
}

func (a *Account) post(strategy auth.Strategy, URI string, body, res interface{}) error {
	return util.POST(a.ctx, a.server, strategy, URI, a.headers, body, res)
}

func (a *Account) patch(strategy auth.Strategy, URI string, body, res interface{}) error {
	return util.PATCH(a.ctx, a.server, strategy, URI, a.headers, body, res)
}

func (a *Account) del(strategy auth.Strategy, URI string) error {
	return util.DELETE(a.ctx, a.server, strategy, URI, a.headers)
}
