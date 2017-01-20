package errors

import (
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/oauth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

// StatusCode gets the status code from an error, defaulting to 500
func StatusCode(err error) int {
	switch t := err.(type) {
	case util.HTTPError:
		return t.Code
	case *util.HTTPError:
		return t.Code
	case *oauth.Error:
		return t.Code
	default:
		return http.StatusInternalServerError
	}
}
