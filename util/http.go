// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/apex/log"
)

var (
	// MaxRedirects specifies the maximum number of redirects an HTTP
	// request should be able to make
	MaxRedirects = 5
)

// HTTPError represents an error coming over HTTP,
// it is not an error with executing the request itself, it is
// an error the server is flaggin to the client.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func (e HTTPError) Error() string {
	return e.Message
}

// checkRedirect implements this clients redirection policy
func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > MaxRedirects {
		return errors.New("Maximum number of redirects reached")
	}

	// use the same headers as before
	req.Header.Set("Authorization", via[len(via)-1].Header.Get("Authorization"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return nil
}

// NewRequest creates a new http.Request that has authorization set up
func newRequest(server, method string, URI string, body io.Reader) (*http.Request, error) {
	URL := fmt.Sprintf("%s%s", server, URI)
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func performRequestBody(ctx log.Interface, server string, strategy auth.Strategy, method, URI string, body interface{}, redirects int) (io.ReadCloser, error) {
	var req *http.Request
	var err error

	if body != nil {
		// body is not nil, so serialize it and pass it in the request
		if err = Validate(body); err != nil {
			return nil, fmt.Errorf("Got an illegal request body: %s", err)
		}

		buf := new(bytes.Buffer)
		encoder := json.NewEncoder(buf)
		err = encoder.Encode(body)
		if err != nil {
			return nil, err
		}
		req, err = newRequest(server, method, URI, buf)
	} else {
		// body is nil so create a nil request
		req, err = newRequest(server, method, URI, nil)
	}

	// decorate the request
	strategy.DecorateRequest(req)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		CheckRedirect: checkRedirect,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// catch deprecated api
	if resp.StatusCode == 410 {
		return nil, fmt.Errorf("API deprecated by The Things Network account server, please update your client")
	} else if warning := resp.Header.Get("Warning"); warning != "" {
		// Warning header has format: 123 - "Message"
		code := warning[0:3]
		message := warning[7 : len(warning)-1]
		ctx.WithFields(log.Fields{
			"code":    code,
			"message": message,
		}).Warn("Got server warning. Make sure the client is up to date.")
	}

	if resp.StatusCode >= 400 {
		var herr HTTPError
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&herr); err != nil {
			// could not decode body as error, just return http error
			return nil, HTTPError{
				Code:    resp.StatusCode,
				Message: resp.Status[4:],
			}
		}

		// fill in blank code
		if herr.Code == 0 {
			herr.Code = resp.StatusCode
		}

		// fill in blank message
		if herr.Message == "" {
			herr.Message = resp.Status[4:]
		}

		return nil, herr
	}

	if resp.StatusCode == 307 {
		if redirects > 0 {
			location := resp.Header.Get("Location")
			return performRequestBody(ctx, server, strategy, method, location, body, redirects-1)
		}
		return nil, fmt.Errorf("Reached maximum number of redirects")
	}

	if resp.StatusCode >= 300 {
		// 307 is handled, 301, 302, 304 cannot be
		return nil, fmt.Errorf("Unexpected %v redirection to %s", resp.StatusCode, resp.Header.Get("Location"))
	}

	return resp.Body, nil
}

// performRequest performs a request and decodes the result
func performRequest(ctx log.Interface, server string, strategy auth.Strategy, method, URI string, rbody, res interface{}, redirects int) error {
	body, err := performRequestBody(ctx, server, strategy, method, URI, rbody, redirects)
	if err != nil {
		return err
	}

	if res != nil {
		defer body.Close()
		decoder := json.NewDecoder(body)
		if err := decoder.Decode(res); err != nil {
			return err
		}

		if err := Validate(res); err != nil {
			return fmt.Errorf("Got an illegal response from server: %s", err)
		}
	}

	return nil
}

// GET does a get request to the account server,  decoding the result into the object pointed to byres
func GET(ctx log.Interface, server string, strategy auth.Strategy, URI string, res interface{}) error {
	return performRequest(ctx, server, strategy, "GET", URI, nil, res, MaxRedirects)
}

// GET does a get request to the account server,  decoding the result into the object pointed to byres
func GETBody(ctx log.Interface, server string, strategy auth.Strategy, URI string) (io.ReadCloser, error) {
	return performRequestBody(ctx, server, strategy, "GET", URI, nil, MaxRedirects)
}

// DELETE does a delete request to the account server
func DELETE(ctx log.Interface, server string, strategy auth.Strategy, URI string) error {
	return performRequest(ctx, server, strategy, "DELETE", URI, nil, nil, MaxRedirects)
}

// POST creates an HTTP Post request to the specified server, with the body
// encoded as JSON, decoding the result into the object pointed to byres
func POST(ctx log.Interface, server string, strategy auth.Strategy, URI string, body, res interface{}) error {
	return performRequest(ctx, server, strategy, "POST", URI, body, res, MaxRedirects)
}

// PUT creates an HTTP Put request to the specified server, with the body
// encoded as JSON, decoding the result into the object pointed to byres
func PUT(ctx log.Interface, server string, strategy auth.Strategy, URI string, body, res interface{}) error {
	return performRequest(ctx, server, strategy, "PUT", URI, body, res, MaxRedirects)
}

// PATCH creates an HTTP Patch request to the specified server, with the body
// encoded as JSON, decoding the result into the object pointed to byres
func PATCH(ctx log.Interface, server string, strategy auth.Strategy, URI string, body, res interface{}) error {
	return performRequest(ctx, server, strategy, "PATCH", URI, body, res, MaxRedirects)
}
