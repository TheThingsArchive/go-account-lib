// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package util

import (
	"net/http"
	"strings"

	"github.com/TheThingsNetwork/go-utils/log"
)

// RoundTripper is a http.RoundTripper that check the Warning header
// and can add extra headers to the request
type RoundTripper struct {
	ctx       log.Interface
	transport http.RoundTripper
	headers   map[string]string
}

// NewRoundTripper creates a new RoundTripper that will add the provided headers
// to each request
func NewRoundTripper(ctx log.Interface, headers map[string]string) *RoundTripper {
	return &RoundTripper{
		ctx:       ctx,
		transport: http.DefaultTransport,
		headers:   headers,
	}
}

func (t *RoundTripper) addHeaders(req *http.Request) {
	for name, value := range t.headers {
		switch strings.ToLower(name) {
		case "authorization":
		default:
			req.Header.Set(name, value)
		}
	}
}

// RoundTrip performs one HTTP roundtrip, adding the headers and checking
// warnings
func (t *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	t.addHeaders(req)
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// get server warnings
	if warning := res.Header.Get("Warning"); len(warning) >= 8 {
		// Warning header has format: 123 - "Message"
		code := warning[0:3]
		message := warning[7 : len(warning)-1]
		if t.ctx != nil {
			t.ctx.WithFields(map[string]interface{}{
				"code":    code,
				"message": message,
			}).Warn("Got server warning. Make sure the client is up to date.")
		}
	}

	return res, nil
}
