// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokenkey

import (
	"testing"

	"github.com/TheThingsNetwork/go-account-lib/tokenkey/cache"
	. "github.com/smartystreets/assertions"
)

var (
	badID   = "bad"
	testID  = "test"
	testURL = "https://test"
	key     = []byte{0x01, 0x02}
	servers = map[string]string{
		testID: testURL,
	}
)

type constCache struct {
	data []byte
}

func (c *constCache) Get(server string) ([]byte, error) {
	return c.data, nil
}

func (c *constCache) Set(server string, data []byte) error {
	return nil
}

func TestHTTPProviderBadServer(t *testing.T) {
	a := New(t)

	provider := HTTPProvider(servers, cache.EmptyCache)

	_, err := provider.Get(badID, false)
	a.So(err, ShouldNotBeNil)

}

func TestHTTPProviderBadKey(t *testing.T) {
	a := New(t)

	cache := &constCache{data: key}
	provider := HTTPProvider(servers, cache)

	_, err := provider.Get(testID, false)
	a.So(err, ShouldNotBeNil)
}
