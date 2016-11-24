// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import jwt "github.com/dgrijalva/jwt-go"

type TTNClaims interface {
	jwt.Claims
	Issuer() string
}

// Claims represents all the claims an access token can have
type Claims struct {
	jwt.StandardClaims
	Client     string              `json:"client"`
	Scope      []string            `json:"scope"`
	Type       string              `json:"type,omitempty"`
	Apps       map[string][]string `json:"apps,omitempty"`
	Gateways   map[string][]string `json:"gateways,omitempty"`
	Components map[string][]string `json:"components,omitempty"`
	Username   string              `json:"username"`
	Email      string              `json:"email"`
	Name       struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
}

func (c *Claims) Issuer() string {
	return c.StandardClaims.Issuer
}

// GatewayClaims represents all the claims an access token can have
type GatewayClaims struct {
	jwt.StandardClaims
	Type           string `json:"type"`
	LocationPublic bool   `json:"location_public"`
	StatusPublic   bool   `json:"status_public"`
}

func (c *GatewayClaims) Issuer() string {
	return c.StandardClaims.Issuer
}
