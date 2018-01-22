// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/TheThingsNetwork/go-account-lib/scope"
	"github.com/TheThingsNetwork/ttn/core/types"
	"golang.org/x/oauth2"
)

// ListGateways list all gateways
func (a *Account) ListGateways() (gateways []Gateway, err error) {
	err = a.get(a.auth, "/api/v2/gateways", &gateways)
	return gateways, err
}

// GatewayStream is a stream of gateways that can be closed
type GatewayStream struct {
	body    io.ReadCloser
	decoder *json.Decoder
}

// Close closes the gateway stream
func (s *GatewayStream) Close() error {
	return s.body.Close()
}

// Next requests the next gateway on the stream, blocking until there is one
// If there are no more gateways, the error will be io.EOF
func (s *GatewayStream) Next() (*Gateway, error) {
	var res gatewayRes
	if s.decoder.More() {
		err := s.decoder.Decode(&res)
		gateway := res.ToGateway()
		return &gateway, err
	}

	// parse last token
	_, err := s.decoder.Token()
	if err != nil {
		return nil, err
	}

	return nil, io.EOF
}

// StreamGateways lists all gateways in a streaming fashion
func (a *Account) StreamGateways() (*GatewayStream, error) {
	body, err := a.gets(a.auth, "/api/v2/gateways")
	if err != nil {
		return nil, err
	}

	stream := &GatewayStream{
		decoder: json.NewDecoder(body),
		body:    body,
	}

	// parse the first array bracket
	_, err = stream.decoder.Token()
	if err != nil {
		return nil, err
	}

	return stream, err
}

type gatewayRes struct {
	Gateway
	GWToken *tokenRes `json:"token,omitempty"`
}

func (r *gatewayRes) ToGateway() Gateway {
	gateway := r.Gateway
	if r.GWToken != nil {
		gateway.Token = r.GWToken.Token()
	}
	return gateway
}

// FindGateway returns the information about a specific gateay
func (a *Account) FindGateway(gatewayID string) (gateway Gateway, err error) {
	var res gatewayRes
	err = a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), &res)
	if err != nil {
		return gateway, err
	}
	return res.ToGateway(), err
}

type registerGatewayReq struct {
	// ID is the ID of the new gateway (required)
	ID string `json:"id"`

	// Country is the country code of the new gateway (required)
	FrequencyPlan string `json:"frequency_plan"`

	// AntennaLocation is the location of the gateway antenna
	AntennaLocation *Location `json:"antenna_location,omitempty"`

	// Attributes is a free-form map of attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// Router is the address (hostname:port) of the router this gateway talks to
	Router string `json:"router,omitempty"`
}

// GatewaySettings represents settings that can be changed on a gateway
type GatewaySettings struct {
	// AntennaLocation is the location of the gateway antenna
	AntennaLocation *Location `json:"location,omitempty"`

	// Attributes is a free-form map of attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// Router is the primary router this gateway talks to
	Router string `json:"router,omitempty"`

	// FallbackRouters are the id's of routers the gateway can connect to when the
	// primary router goes down
	FallbackRouters []string `json:"fallback_routers,omitempty"`
}

// RegisterGateway registers a new gateway on the account server
func (a *Account) RegisterGateway(gatewayID string, frequencyPlan string, opts GatewaySettings) (gateway Gateway, err error) {
	if gatewayID == "" {
		return gateway, errors.New("Cannot create gateway: no ID given")
	}

	if frequencyPlan == "" {
		return gateway, errors.New("Cannot create gateway: no frequency plan given")
	}

	req := registerGatewayReq{
		ID:              gatewayID,
		FrequencyPlan:   frequencyPlan,
		AntennaLocation: opts.AntennaLocation,
		Attributes:      opts.Attributes,
		Router:          opts.Router,
	}

	err = a.post(a.auth, "/api/v2/gateways", req, &gateway)
	return gateway, err
}

// GetGatewayToken gets the gateway token
func (a *Account) GetGatewayToken(gatewayID string) (*oauth2.Token, error) {
	token := new(tokenRes)
	err := a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/token", gatewayID), token)
	if err != nil {
		return nil, err
	}

	return token.Token(), nil
}

// DeleteGateway removes a gateway from the account server
func (a *Account) DeleteGateway(gatewayID string) error {
	return a.del(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID))
}

// GrantGatewayRights grants rights to a collaborator of the gateway
func (a *Account) GrantGatewayRights(gatewayID string, username string, rights []types.Right) error {
	req := grantReq{
		Rights: rights,
	}
	return a.put(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/collaborators/%s", gatewayID, username), req, nil)
}

// RetractGatewayRights removes rights from a collaborator of the gateway
func (a *Account) RetractGatewayRights(gatewayID string, username string) error {
	return a.del(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/collaborators/%s", gatewayID, username))
}

// GatewayEdits contains editable fields of gateways
type GatewayEdits struct {
	Owner           string             `json:"owner,omitempty"`
	LocationPublic  *bool              `json:"location_public,omitempty"`
	StatusPublic    *bool              `json:"status_public,omitempty"`
	OwnerPublic     *bool              `json:"owner_public,omitempty"`
	FrequencyPlan   string             `json:"frequency_plan,omitempty"`
	AutoUpdate      *bool              `json:"auto_update,omitempty"`
	AntennaLocation *Location          `json:"antenna_location,omitempty"`
	Attributes      *GatewayAttributes `json:"attributes,omitempty"`
	Router          *string            `json:"router,omitempty"`
	FallbackRouters *[]string          `json:"fallback_routers,omitempty"`
	BetaUpdates     *bool              `json:"beta_updates,omitempty"`
}

// EditGateway edits the fields of a gateway
func (a *Account) EditGateway(gatewayID string, edits GatewayEdits) error {
	return a.patch(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), edits, nil)
}

// TransferOwnership transfers the owenership of the gateway to another user
func (a *Account) TransferOwnership(gatewayID, username string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Owner: username,
	})
}

// ChangeFrequencyPlan changes the requency plan of a gateway
func (a *Account) ChangeFrequencyPlan(gatewayID, plan string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		FrequencyPlan: plan,
	})
}

// ChangeLocation changes the location of the gateway, set to nil, nil if you
// want to remove the location
func (a *Account) ChangeLocation(gatewayID string, latitude, longitude float64) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		AntennaLocation: &Location{
			Longitude: longitude,
			Latitude:  latitude,
		},
	})
}

// ChangeAltitude changes the altitude of the gateway with the specified ID
func (a *Account) ChangeAltitude(gatewayID string, altitude int) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		AntennaLocation: &Location{
			Altitude: altitude,
		},
	})
}

// ChangeRouter changes the router the gateway talks to
func (a *Account) ChangeRouter(gatewayID string, router string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Router: &router,
	})
}

// GatewayRights returns the rights the current account client has to a certain
// gateway
func (a *Account) GatewayRights(gatewayID string) (rights []types.Right, err error) {
	err = a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/rights", gatewayID), &rights)
	if err != nil {
		return nil, err
	}

	return rights, nil
}

// GatewayCollaborators fetches the gateway collaborators
func (a *Account) GatewayCollaborators(gwID string) ([]Collaborator, error) {
	collaborators := make([]Collaborator, 0)
	err := a.get(a.auth.WithScope(scope.Gateway(gwID)), fmt.Sprintf("/api/v2/gateways/%s/collaborators", gwID), &collaborators)
	if err != nil {
		return nil, err
	}

	return collaborators, nil
}
