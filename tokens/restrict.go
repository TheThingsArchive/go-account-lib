// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type restrictRequest struct {
	Scope []string `json:"scope"`
}

type restrictResponse struct {
	AccessToken string `json:"access_token"`
}

// HTTPError is an error that arises from HTTP
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func (e HTTPError) Error() string {
	return e.Message
}

// RestrictScope requests a new token with a different, more specific scope
func RestrictScope(server string, token string, scope []string) (string, error) {
	URL := fmt.Sprintf("%s/users/restrict-token", server)

	body := restrictRequest{
		Scope: scope,
	}
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", URL, buf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		var herr HTTPError
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&herr); err != nil {
			// could not decode body as error, just return http error
			return "", HTTPError{
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

		return "", herr
	}

	var res restrictResponse
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return "", err
	}

	return res.AccessToken, nil
}
