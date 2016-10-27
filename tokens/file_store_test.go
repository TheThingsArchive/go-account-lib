// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package tokens

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
)

func lines(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(strings.Split(string(data), "\n")), nil
}

func TestFileStore(t *testing.T) {
	a := New(t)

	dir, err := ioutil.TempDir("", "test")
	a.So(err, ShouldBeNil)
	defer os.RemoveAll(dir)

	fname := path.Join(dir, "store.tokens")
	store := FileStore(fname)

	tok, err := store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldBeEmpty)

	err = store.Set(parent, scopes, token, time.Hour)
	a.So(err, ShouldBeNil)

	tok, err = store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, token)

	err = store.Set(otherParent, scopes, token, time.Hour)
	a.So(err, ShouldBeNil)

	n, err := lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 2)

	tok, err = store.Get(otherParent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, token)

	err = store.Set(otherParent, scopes, token, -time.Hour)
	a.So(err, ShouldBeNil)

	n, err = lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 1)

	tok, err = store.Get(otherParent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldBeEmpty)

	err = store.Set(parent, scopes, otherToken, time.Hour)
	a.So(err, ShouldBeNil)

	n, err = lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 1)

	tok, err = store.Get(parent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, otherToken)
}
