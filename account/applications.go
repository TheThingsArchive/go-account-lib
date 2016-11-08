// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/TheThingsNetwork/go-account-lib/scope"
	"github.com/TheThingsNetwork/ttn/core/types"
)

// ListApplications list all applications
func (a *Account) ListApplications() (apps []Application, err error) {
	err = a.get(a.auth.WithScope(scope.Apps), "/api/v2/applications", &apps)
	return apps, err
}

// GatewayStream is a stream of gateways that can be closed
type ApplicationStream struct {
	body    io.ReadCloser
	decoder *json.Decoder
}

// Close closes the gateway stream
func (s *ApplicationStream) Close() error {
	return s.body.Close()
}

// Next requests the next gateway on the stream, blocking until there is one
// If there are no more gateways, the error will be io.EOF
func (s *ApplicationStream) Next() (*Application, error) {
	var application Application
	if s.decoder.More() {
		err := s.decoder.Decode(&application)
		return &application, err
	}

	// parse last token
	_, err := s.decoder.Token()
	if err != nil {
		return nil, err
	}

	return nil, io.EOF
}

// StreamGateways lists all gateways in a streaming fashion
func (a *Account) StreamApplications() (*ApplicationStream, error) {
	body, err := a.gets(a.auth, "/api/v2/applications")
	if err != nil {
		return nil, err
	}

	stream := &ApplicationStream{
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

// FindApplication gets a specific application from the account server
func (a *Account) FindApplication(appID string) (app Application, err error) {
	err = a.get(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s", appID), &app)
	return app, err
}

type createApplicationReq struct {
	Name  string         `json:"name" valid:"required"`
	AppID string         `json:"id"   valid:"required"`
	EUIs  []types.AppEUI `json:"euis"`
}

// CreateApplication creates a new application on the account server
func (a *Account) CreateApplication(appID string, name string, EUIs []types.AppEUI) (app Application, err error) {
	body := createApplicationReq{
		Name:  name,
		AppID: appID,
		EUIs:  EUIs,
	}

	err = a.post(a.auth, "/api/v2/applications", &body, &app)
	return app, err
}

// DeleteApplication deletes an application
func (a *Account) DeleteApplication(appID string) error {
	return a.del(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s", appID))
}

type grantReq struct {
	Rights []types.Right `json:"rights"`
}

// Grant adds a collaborator to the application
func (a *Account) Grant(appID string, username string, rights []types.Right) error {
	req := grantReq{
		Rights: rights,
	}
	return a.put(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/collaborators/%s", appID, username), req, nil)
}

// Retract removes rights from a collaborator of the application
func (a *Account) Retract(appID string, username string) error {
	return a.del(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/collaborators/%s", appID, username))
}

type addAccessKeyReq struct {
	Name   string        `json:"name"   valid:"required"`
	Rights []types.Right `json:"rights" valid:"required"`
}

// AddAccessKey adds an access key to the application with the specified name
// and rights
func (a *Account) AddAccessKey(appID string, name string, rights []types.Right) (key types.AccessKey, err error) {
	body := addAccessKeyReq{
		Name:   name,
		Rights: rights,
	}
	err = a.post(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/access-keys", appID), body, &key)
	return key, err
}

// RemoveAccessKey removes the specified access key from the application
func (a *Account) RemoveAccessKey(appID string, name string) error {
	return a.del(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/access-keys/%s", appID, name))
}

type editAppReq struct {
	Name string `json:"name,omitempty"`
}

// ChangeName changes the application name
func (a *Account) ChangeName(appID string, name string) error {
	body := editAppReq{
		Name: name,
	}
	return a.patch(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s", appID), body, nil)
}

// AddEUI adds an EUI to the applications list of EUIs
func (a *Account) AddEUI(appID string, eui types.AppEUI) error {
	return a.put(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/euis/%s", appID, eui.String()), nil, nil)
}

type genEUIRes struct {
	EUI types.AppEUI `json:"eui"`
}

// GenerateEUI creates a new EUI for the application
func (a *Account) GenerateEUI(appID string) (*types.AppEUI, error) {
	var res genEUIRes
	err := a.post(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/euis", appID), nil, &res)
	if err != nil {
		return nil, err
	}
	return &res.EUI, nil
}

// RemoveEUI removes the specified EUI from the application
func (a *Account) RemoveEUI(appID string, eui types.AppEUI) error {
	return a.del(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/euis/%s", appID, eui.String()))
}

// AppRights returns the rights the current account client has to a certain
// application
func (a *Account) AppRights(appID string) (rights []types.Right, err error) {
	err = a.get(a.auth.WithScope(scope.App(appID)), fmt.Sprintf("/api/v2/applications/%s/rights", appID), &rights)
	if err != nil {
		return nil, err
	}

	return rights, nil
}
