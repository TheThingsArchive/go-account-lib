// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"encoding/json"
	"time"

	"github.com/TheThingsNetwork/ttn/core/types"
	"golang.org/x/oauth2"
)

// Application represents an application on The Things Network
type Application struct {
	ID            string            `json:"id"   valid:"required"`
	Name          string            `json:"name" valid:"required"`
	EUIs          []types.AppEUI    `json:"euis,omitempty"`
	AccessKeys    []types.AccessKey `json:"access_keys,omitempty"`
	Created       time.Time         `json:"created,omitempty"`
	Collaborators []Collaborator    `json:"collaborators,omitempty"`
}

// Collaborator is a user that has rights to a certain application
type Collaborator struct {
	Username string        `json:"username" valid:"required"`
	Rights   []types.Right `json:"rights"   valid:"required"`
}

// HasRight checks if the collaborator has a specific right
func (c *Collaborator) HasRight(right types.Right) bool {
	for _, r := range c.Rights {
		if r == right {
			return true
		}
	}
	return false
}

// Profile represents the profile of a user
type Profile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     *Name  `json:"name"`
}

// Name represents the full name of a user
type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// Component represents a component on the newtork
type Component struct {
	Type    string    `json:"type"`
	ID      string    `json:"id"`
	Created time.Time `json:"created,omitempty"`
}

// String implements the Stringer interface for Name
func (n *Name) String() string {
	return n.First + " " + n.Last
}

type gatewayToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint64 `json:"expires_in"`
}

// Token transfroms a gateway token to oauth token with correct expiry
func (token *gatewayToken) Token() *oauth2.Token {
	if token == nil {
		return nil
	}

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		Expiry:      time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}
}

type Placement string

const (
	Indoor  Placement = "indoor"
	Outdoor Placement = "outdoor"
)

type GatewayAttributes struct {
	Brand        *string    `json:"brand,omitempty"`
	Model        *string    `json:"model,omitempty"`
	Placement    *Placement `json:"placement,omitempty"`
	AntennaType  *string    `json:"antenna_type,omitempty"`
	AntennaModel *string    `json:"antenna_model,omitempty"`
	Description  *string    `json:"description,omitempty"`
}

// Gateway represents a gateway on the account server
type Gateway struct {
	ID               string         `json:"id" valid:"required"`
	Activated        bool           `json:"activated"`
	FrequencyPlan    string         `json:"frequency_plan"`
	FrequencyPlanURL string         `json:"frequency_plan_url"`
	PublicRights     []types.Right  `json:"public_rights"`
	LocationPublic   bool           `json:"location_public"`
	StatusPublic     bool           `json:"status_public"`
	AutoUpdate       bool           `json:"auto_update"`
	Location         *Location      `json:"location"`
	Altitude         float64        `json:"altitude"`
	Collaborators    []Collaborator `json:"collaborators"`
	Key              string         `json:"key"`
	token            *gatewayToken  `json:"token,omitempty"`
	Token            *oauth2.Token
	Attributes       GatewayAttributes `json:"attributes"`
	Router           string            `json:"string"`
}

// Location is the GPS location of a gateway
type Location struct {
	// Empty denotes that the location is not given as oposed to (0, 0) which is a
	// valid location
	Empty     bool    `json:"-"`
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

type FrequencyPlan struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type location struct {
	Lng *float64 `json:"lng,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
}

// UnmarshalJSON is a custom unmarshaller for location that allows falsy types
// false | {} | null -> Location{ Empty: true }
// ... -> Location
func (loc *Location) UnmarshalJSON(b []byte) error {
	loc.Empty = true
	loc.Longitude = 0
	loc.Latitude = 0

	if string(b) == "false" || string(b) == "null" {
		return nil
	}

	var l location
	err := json.Unmarshal(b, &l)
	if err != nil {
		return err
	}

	if l.Lat == nil || l.Lng == nil {
		return nil
	}

	loc.Empty = false
	loc.Latitude = *l.Lat
	loc.Longitude = *l.Lng

	return nil
}

// MarshalJSON is a custom json marshaller for Location that maps an empty
// Location to `false`
func (loc Location) MarshalJSON() ([]byte, error) {
	if loc.Empty {
		return []byte("false"), nil
	}

	l := location{
		Lat: &loc.Latitude,
		Lng: &loc.Longitude,
	}

	return json.Marshal(l)
}
