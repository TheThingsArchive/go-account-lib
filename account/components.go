// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"fmt"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/scope"
)

type ComponentType string

const (
	Handler ComponentType = "handler"
	Router  ComponentType = "router"
	Broker  ComponentType = "broker"
)

func plural(in string) (string, error) {
	switch in {
	case string(Handler):
		return "handlers", nil
	case string(Router):
		return "routers", nil
	case string(Broker):
		return "brokers", nil
	default:
		return "", fmt.Errorf("Invalid component type `%s`", in)
	}
}

// ListComponents lists all of the users components
func (a *Account) ListComponents() ([]Component, error) {
	components := make([]Component, 0)
	err := a.get(a.auth, "/api/v2/components", &components)
	return components, err
}

// FindComponent finds a comonent of the specified type with the specified id
func (a *Account) FindComponent(typ, id string) (component Component, err error) {
	p, err := plural(typ)
	if err != nil {
		return component, err
	}
	err = a.get(a.auth.WithScope(scope.Component(id)), fmt.Sprintf("/api/v2/components/%s/%s", p, id), &component)
	return component, err
}

// FindBroker finds a broker with the specified id
func (a *Account) FindBroker(id string) (component Component, err error) {
	return a.FindComponent("broker", id)
}

// FindRouter finds a router with the specified id
func (a *Account) FindRouter(id string) (component Component, err error) {
	return a.FindComponent("router", id)
}

// FindHandler finds a handler with the specified id
func (a *Account) FindHandler(id string) (component Component, err error) {
	return a.FindComponent("handler", id)
}

type createComponentReq struct {
	ID string `json:"id" valid:"required"`
}

// CreateComponent creates a component with the specified type and id
func (a *Account) CreateComponent(typ, id string) error {
	p, err := plural(typ)
	if err != nil {
		return err
	}

	body := createComponentReq{
		ID: id,
	}
	return a.post(a.auth, fmt.Sprintf("/api/v2/components/%s", p), body, nil)
}

// CreateBroker creates a broker with the specified id
func (a *Account) CreateBroker(id string) error {
	return a.CreateComponent("broker", id)
}

// CreateRouter creates a Router with the specified id
func (a *Account) CreateRouter(id string) error {
	return a.CreateComponent("router", id)
}

// CreateHandler creates a handler with the specified id
func (a *Account) CreateHandler(id string) error {
	return a.CreateComponent("handler", id)
}

type componentTokenRes struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// ComponentToken fetches a token for the component with the given
// type and id
func (a *Account) ComponentToken(typ, id string) (token string, err error) {
	p, err := plural(typ)
	if err != nil {
		return "", err
	}

	var res componentTokenRes
	err = a.get(a.auth.WithScope(scope.Component(id)), fmt.Sprintf("/api/v2/components/%s/%s/token", p, id), &res)
	return res.Token, err
}

// BrokerToken gets the specified brokers token
func (a *Account) BrokerToken(id string) (token string, err error) {
	return a.ComponentToken("broker", id)
}

// RouterToken gets the specified routers token
func (a *Account) RouterToken(id string) (token string, err error) {
	return a.ComponentToken("router", id)
}

// HandlerToken gets the specified handlers token
func (a *Account) HandlerToken(id string) (token string, err error) {
	return a.ComponentToken("handler", id)
}
