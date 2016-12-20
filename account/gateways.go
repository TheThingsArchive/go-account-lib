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
	var gateway Gateway
	if s.decoder.More() {
		err := s.decoder.Decode(&gateway)
		gateway.Token = gateway.token.Token()
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

// FindGateway returns the information about a specific gateay
func (a *Account) FindGateway(gatewayID string) (gateway Gateway, err error) {
	err = a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), &gateway)
	gateway.Token = gateway.token.Token()
	return gateway, err
}

type registerGatewayReq struct {
	// ID is the ID of the new gateway (required)
	ID string `json:"id"`

	// Country is the country code of the new gateway (required)
	FrequencyPlan string `json:"frequency_plan"`

	// Location is the location of the new gateway
	Location *Location `json:"location,omitempty"`

	// Attributes is a free-form map of attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// Router is the router this gateway talks to
	Router string `json:"router,omitempty"`
}

// RegisterGateway registers a new gateway on the account server
func (a *Account) RegisterGateway(gatewayID string, frequencyPlan string, location *Location) (gateway Gateway, err error) {
	if gatewayID == "" {
		return gateway, errors.New("Cannot create gateway: no ID given")
	}

	if frequencyPlan == "" {
		return gateway, errors.New("Cannot create gateway: no FrequencyPlan given")
	}

	req := registerGatewayReq{
		ID:            gatewayID,
		FrequencyPlan: frequencyPlan,
		Location:      location,
	}

	err = a.post(a.auth, "/api/v2/gateways", req, &gateway)
	return gateway, err
}

// GetGatewayToken gets the gateway token
func (a *Account) GetGatewayToken(gatewayID string) (*oauth2.Token, error) {
	token := new(gatewayToken)
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
	Owner         string             `json:"owner,omitempty"`
	PublicRights  *[]types.Right     `json:"public_rights,omitempty"`
	FrequencyPlan string             `json:"frequency_plan,omitempty"`
	AutoUpdate    *bool              `json:"auto_update,omitempty"`
	Location      *Location          `json:"location,omitempty"`
	Altitude      float64            `json:"altitude,omitempty"`
	Attributes    *GatewayAttributes `json:"attributes,omitempty"`
	Router        string             `json:"router,omitempty"`
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

// SetPublicRights changes the publicily visible rights of the gateway
func (a *Account) SetPublicRights(gatewayID string, rights []types.Right) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		PublicRights: &rights,
	})
}

// ChangeFrequencyPlan changes the requency plan of a gateway
func (a *Account) ChangeFrequencyPlan(gatewayID, plan string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		FrequencyPlan: plan,
	})
}

// ChangeLocation changes the location of the gateway
func (a *Account) ChangeLocation(gatewayID string, latitude, longitude float64) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Location: &Location{
			Longitude: longitude,
			Latitude:  latitude,
		},
	})
}

// ChangeAltitude changes the altitude of the gateway with the specified ID
func (a *Account) ChangeAltitude(gatewayID string, altitude float64) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Altitude: altitude,
	})
}

// ChangeRouter changes the router the gateway talks to
func (a *Account) ChangeRouter(gatewayID string, router string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Router: router,
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
