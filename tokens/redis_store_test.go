// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
	redis "gopkg.in/redis.v3"
)

func getRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: "",
		DB:       1,
	})
}

func TestRedisStore(t *testing.T) {
	a := New(t)

	client := getRedisClient()
	store := RedisStore(client)

	// getting from an empty store should work
	res, err := store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")

	// setting to a new store should work
	err = store.Set(parent, scopes, token, time.Second)
	a.So(err, ShouldBeNil)

	// getting from a not-so-new store should still work
	res, err = store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, token)

	// trying to get an expired token should not work
	err = store.Set(parent, []string{"scope2"}, token, time.Duration(-1))
	a.So(err, ShouldBeNil)

	res, err = store.Get(parent, otherScope)
	a.So(err, ShouldBeNil)
	a.So(res, ShouldEqual, "")
}
