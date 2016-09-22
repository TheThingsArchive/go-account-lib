// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package auth

import "net/http"

type public struct{}

func (p *public) DecorateRequest(req *http.Request) {
}

var Public = &public{}
