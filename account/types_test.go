// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"encoding/json"
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

func TestMarshalLocation(t *testing.T) {
	a := assertions.New(t)

	// marshal empty location
	{
		var empty *AntennaLocation

		str, err := json.Marshal(empty)
		a.So(err, assertions.ShouldBeNil)
		a.So(string(str), assertions.ShouldEqual, "null")
	}

	// 	// marshal non-emty location
	// 	{
	// 		empty := Location{
	// 			Longitude: float64(10),
	// 			Latitude:  float64(33),
	// 		}
	//
	// 		str, err := json.Marshal(empty)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(string(str), assertions.ShouldEqual, `{"lng":10,"lat":33}`)
	// 	}
	//
	// 	// marshal non-emty locatio (0,0)n
	// 	{
	// 		empty := Location{
	// 			Longitude: float64(0),
	// 			Latitude:  float64(0),
	// 		}
	//
	// 		str, err := json.Marshal(empty)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(string(str), assertions.ShouldEqual, `{"lng":0,"lat":0}`)
	// 	}
	//
	// 	// marshal nested location
	// 	{
	// 		// test nested location
	// 		foo := struct {
	// 			Location Location `json:"location"`
	// 		}{
	// 			Location: Location{Empty: true},
	// 		}
	//
	// 		str, err := json.Marshal(foo)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(string(str), assertions.ShouldEqual, `{"location":false}`)
	// 	}
	// }
	//
	// func TestUnmarshalLocation(t *testing.T) {
	// 	a := assertions.New(t)
	// 	var loc Location
	//
	// 	// unmarshal empty value
	// 	{
	// 		str := `{}`
	// 		err := json.Unmarshal([]byte(str), &loc)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(loc.Empty, assertions.ShouldBeTrue)
	// 		a.So(loc.Longitude, assertions.ShouldEqual, float64(0))
	// 		a.So(loc.Latitude, assertions.ShouldEqual, float64(0))
	// 	}
	//
	// 	// unmarshal false
	// 	{
	// 		str := `false`
	// 		err := json.Unmarshal([]byte(str), &loc)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(loc.Empty, assertions.ShouldBeTrue)
	// 		a.So(loc.Longitude, assertions.ShouldEqual, float64(0))
	// 		a.So(loc.Latitude, assertions.ShouldEqual, float64(0))
	// 	}
	//
	// 	// umarshal non empty value
	// 	{
	// 		str := `{"lat":33, "lng":10}`
	// 		err := json.Unmarshal([]byte(str), &loc)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(loc.Empty, assertions.ShouldBeFalse)
	// 		a.So(loc.Longitude, assertions.ShouldEqual, float64(10))
	// 		a.So(loc.Latitude, assertions.ShouldEqual, float64(33))
	// 	}
	//
	// 	// unmarshal nested location
	// 	{
	// 		// test nested location false
	// 		var foo struct {
	// 			Location Location `json:"location"`
	// 		}
	//
	// 		str := `{ "location": false }`
	// 		err := json.Unmarshal([]byte(str), &foo)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(foo.Location.Empty, assertions.ShouldBeTrue)
	// 		a.So(foo.Location.Longitude, assertions.ShouldEqual, float64(0))
	// 		a.So(foo.Location.Latitude, assertions.ShouldEqual, float64(0))
	// 	}
	//
	// 	// unmarshal nested location {}
	// 	{
	// 		// test nested location
	// 		var foo struct {
	// 			Location Location `json:"location"`
	// 		}
	//
	// 		str := `{ "location": {} }`
	// 		err := json.Unmarshal([]byte(str), &foo)
	// 		a.So(err, assertions.ShouldBeNil)
	// 		a.So(foo.Location.Empty, assertions.ShouldBeTrue)
	// 		a.So(foo.Location.Longitude, assertions.ShouldEqual, float64(0))
	// 		a.So(foo.Location.Latitude, assertions.ShouldEqual, float64(0))
	// 	}

}
