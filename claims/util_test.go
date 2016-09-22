// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"testing"

	. "github.com/smartystreets/assertions"
)

const (
	fst  = "first"
	snd  = "second"
	thd  = "third"
	diff = "different"
)

var container = []string{fst, snd, thd}

func TestContains(t *testing.T) {
	a := New(t)

	a.So(contains(container, fst), ShouldBeTrue)
	a.So(contains(container, snd), ShouldBeTrue)
	a.So(contains(container, thd), ShouldBeTrue)
	a.So(contains(container, diff), ShouldBeFalse)
}
