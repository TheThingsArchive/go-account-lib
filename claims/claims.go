// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import jwt "github.com/dgrijalva/jwt-go"

// TTNClaims is the interface all claims
// that are issued by TTN need to adhere to
type ClaimsWithIssuer interface {
	jwt.Claims
	GetIssuer() string
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

// Issuer returns the issuer of the claims
func (c *Claims) GetIssuer() string {
	return c.StandardClaims.Issuer
}

// GatewayClaims represents all the claims an access token can have
type GatewayClaims struct {
	jwt.StandardClaims
	Type           string `json:"type"`
	LocationPublic bool   `json:"location_public"`
	StatusPublic   bool   `json:"status_public"`
}

// Issuer returns the issuer of the claims
func (c *GatewayClaims) GetIssuer() string {
	return c.StandardClaims.Issuer
}

// ComponentClaims represent the claims a network component can have
type ComponentClaims struct {
	jwt.StandardClaims
	Type string `json:"type"`
}

// GetIssuer returns the issuer of the claims
func (c *ComponentClaims) GetIssuer() string {
	return c.StandardClaims.Issuer
}
