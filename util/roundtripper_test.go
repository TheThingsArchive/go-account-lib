package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/assertions"
)

const (
	uri = "/foo"
)

func HeadersHandler(a *Assertion, method string, headers map[string]string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.So(r.RequestURI, ShouldEqual, uri)
		a.So(r.Method, ShouldEqual, method)

		resp := struct{}{}

		for name, val := range headers {
			a.So(r.Header.Get(name), ShouldEqual, val)
		}
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(&resp)
	})
}

func TestGet(t *testing.T) {
	a := New(t)

	// nil headers
	{
		var headers map[string]string
		client := &http.Client{
			Transport: NewRoundTripper(nil, headers),
		}
		server := httptest.NewServer(HeadersHandler(a, "GET", headers))
		defer server.Close()

		_, err := client.Get(fmt.Sprintf("%s%s", server.URL, uri))
		a.So(err, ShouldBeNil)
	}

	// nil headers
	{
		headers := map[string]string{
			"Foo": "bar",
		}

		client := &http.Client{
			Transport: NewRoundTripper(nil, headers),
		}
		server := httptest.NewServer(HeadersHandler(a, "GET", headers))
		defer server.Close()

		_, err := client.Get(fmt.Sprintf("%s%s", server.URL, uri))
		a.So(err, ShouldBeNil)
	}
}
