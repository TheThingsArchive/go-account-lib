package errors

import (
	"net/http"

	"github.com/TheThingsNetwork/go-account-lib/oauth"
	"github.com/TheThingsNetwork/go-account-lib/util"
)

// Get a status code from an error
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
