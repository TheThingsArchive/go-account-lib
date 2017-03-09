// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"testing"

	"github.com/TheThingsNetwork/ttn/core/types"
	"github.com/smartystreets/assertions"
)

func TestCollaboratorRights(t *testing.T) {
	a := assertions.New(t)
	c := Collaborator{
		Username: "username",
		Rights: []types.Right{
			types.Right("right"),
		},
	}

	a.So(c.HasRight(types.Right("right")), assertions.ShouldBeTrue)
	a.So(c.HasRight(types.Right("foo")), assertions.ShouldBeFalse)
}

func TestNameString(t *testing.T) {
	a := assertions.New(t)
	name := Name{
		First: "John",
		Last:  "Doe",
	}

	a.So(name.String(), assertions.ShouldEqual, "John Doe")
}
