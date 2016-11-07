// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"encoding/json"
	"regexp"
	"strconv"
)

type OAuthError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type response struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

func (e *OAuthError) Error() string {
	return e.Description
}

var codeRe = regexp.MustCompile("oauth2: cannot fetch token: (\\d+)")
var respRe = regexp.MustCompile("Response: ({.*})")

func fromError(err error) *OAuthError {
	if err == nil {
		return nil
	}

	oerr := &OAuthError{
		Code:        500,
		Description: err.Error(),
	}

	str := err.Error()

	matches := codeRe.FindStringSubmatch(str)
	if len(matches) < 2 {
		return oerr
	}

	code, err := strconv.Atoi(matches[1])
	if err != nil {
		return oerr
	}

	oerr.Code = code

	matches = respRe.FindStringSubmatch(str)
	if len(matches) < 2 {
		return oerr
	}

	var resp response
	err = json.Unmarshal([]byte(matches[1]), &resp)
	if err != nil {
		return oerr
	}

	if resp.Error != "" {
		oerr.Description = resp.Error
	}

	if resp.Description != "" {
		oerr.Description = resp.Description
	}

	if oerr.Description == "" {
		oerr.Description = str
	}

	return oerr
}
