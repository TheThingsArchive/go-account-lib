// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package oauth

import (
	"encoding/json"
	"regexp"
	"strconv"
)

// Error is an error that can occure during an OAuth exchange
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type response struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

// Error returns the error description
func (e *Error) Error() string {
	return e.Description
}

var codeRe = regexp.MustCompile("oauth2: cannot fetch token: (\\d+)")
var respRe = regexp.MustCompile("Response: ({.*})")

func fromError(orig error) error {
	if orig == nil {
		return orig
	}

	str := orig.Error()

	oerr := &Error{
		Code:        500,
		Description: str,
	}

	matches := codeRe.FindStringSubmatch(str)
	if len(matches) < 2 {
		return orig
	}

	code, err := strconv.Atoi(matches[1])
	if err != nil {
		return orig
	}

	oerr.Code = code

	matches = respRe.FindStringSubmatch(str)
	if len(matches) < 2 {
		return orig
	}

	var resp response
	err = json.Unmarshal([]byte(matches[1]), &resp)
	if err != nil {
		return orig
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
