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

const (
	testParent  = "thisistheparenttoken"
	testParent2 = "thisisanotherparent"
	testToken   = "thisisthetoken"
	testToken2  = "thisisanothertoken"
)

var (
	testScopes = []string{testScope}
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

	tok, err := store.Get(testParent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldBeEmpty)

	err = store.Set(testParent, testScopes, testToken, time.Hour)
	a.So(err, ShouldBeNil)

	tok, err = store.Get(testParent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, testToken)

	err = store.Set(testParent2, testScopes, testToken, time.Hour)
	a.So(err, ShouldBeNil)

	n, err := lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 2)

	tok, err = store.Get(testParent2, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, testToken)

	err = store.Set(testParent2, testScopes, testToken, -time.Hour)
	a.So(err, ShouldBeNil)

	n, err = lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 1)

	tok, err = store.Get(testParent2, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldBeEmpty)

	err = store.Set(testParent, testScopes, testToken2, time.Hour)
	a.So(err, ShouldBeNil)

	n, err = lines(fname)
	a.So(err, ShouldBeNil)
	a.So(n, ShouldEqual, 1)

	tok, err = store.Get(testParent, testScope)
	a.So(err, ShouldBeNil)
	a.So(tok, ShouldEqual, testToken2)
}
