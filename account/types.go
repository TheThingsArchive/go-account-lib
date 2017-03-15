// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
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

type tokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint64 `json:"expires_in"`
}

// Token transfroms a gateway token to oauth token with correct expiry
func (token *tokenRes) Token() *oauth2.Token {
	if token == nil {
		return nil
	}

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		Expiry:      time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}
}

// Placement is represents antenna placement
type Placement string

const (
	// Indoor represents an indoor gateway
	Indoor Placement = "indoor"

	// Outdoor represents an outdoor gateway
	Outdoor Placement = "outdoor"
)

// GatewayAttributes represents attributes of a gateway
type GatewayAttributes struct {
	Brand        *string    `json:"brand,omitempty"`
	Model        *string    `json:"model,omitempty"`
	Placement    *Placement `json:"placement,omitempty"`
	AntennaType  *string    `json:"antenna_type,omitempty"`
	AntennaModel *string    `json:"antenna_model,omitempty"`
	Description  *string    `json:"description,omitempty"`
}

// GatewayRouter is the description of a router that the gateway should connect to
type GatewayRouter struct {
	ID          string `json:"id"`
	NetAddress  string `json:"address,omitempty"`
	MQTTAddress string `json:"mqtt_address,omitempty"`
}

// Gateway represents a gateway on the account server
type Gateway struct {
	ID               string            `json:"id" valid:"required"`
	Activated        bool              `json:"activated"`
	FrequencyPlan    string            `json:"frequency_plan"`
	FrequencyPlanURL string            `json:"frequency_plan_url"`
	PublicRights     []types.Right     `json:"public_rights"`
	AutoUpdate       bool              `json:"auto_update"`
	AntennaLocation  *AntennaLocation  `json:"antenna_location"`
	Collaborators    []Collaborator    `json:"collaborators"`
	Key              string            `json:"key"`
	Token            *oauth2.Token     `json:"token,omitempty"`
	Attributes       GatewayAttributes `json:"attributes"`
	Router           *GatewayRouter    `json:"router"`
	FallbackRouters  []GatewayRouter   `json:"fallback_routers"`
	Owner            struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"owner"`
}

// IsPublic checks whether or not the given right is public on the gateway
func (gtw *Gateway) IsPublic(right types.Right) bool {
	for _, r := range gtw.PublicRights {
		if r == right {
			return true
		}
	}

	return false
}

// AntennaLocation is the GPS location of a gateway antenna
type AntennaLocation struct {
	// Longitude is the GPS longitude of the gateway antenna, it can be empty if
	// the location is not set
	Longitude *float64 `json:"longitude,omitempty"`

	// Latitude is the GPS latitude of the gateway antenna, it can be empty if
	// the location is not set
	Latitude *float64 `json:"latitude,omitempty"`

	// Altitude is the height of the gateway antenna (with respect to sea level)
	Altitude *float64 `json:"altitude,omitempty"`
}

// FrequencyPlan is the frequency plan used by a gateway
type FrequencyPlan struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	URL           string `json:"url"`
	BaseFrequency int    `json:"base_freq"`
}
