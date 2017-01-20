package errors

import (
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/oauth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

// StatusCode gets the status code from an error, defaulting to 500
func StatusCode(err error) int {
	switch err.(type) {
	case util.HTTPError:
		return err.Code
	case oauth.Error:
		return err.Code
	default:
		return http.StatusInternalServerError
	}
}
