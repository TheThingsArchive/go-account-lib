package account

import (
	"fmt"

	"github.com/TheThingsNetwork/go-account-lib/scope"
)

// OAuthClient represents a user-created OAuth client
type OAuthClient struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Secret      string   `json:"secret,omitempty"`
	URI         string   `json:"uri,omitempty"`
	Grants      []Grant  `json:"grants"`
	Scope       []string `json:"scope"`
	Accepted    bool     `json:"accepted"`
	Rejected    bool     `json:"rejected"`
	Official    bool     `json:"official"`
}

// Grant is an OAuth grant
type Grant string

const (
	// AuthorizationCode represents the OAuth 2.0 authorization code grant
	AuthorizationCode Grant = "authorization_code"

	// RefreshToken represents the OAuth 2.0 refresh token grant
	RefreshToken Grant = "refresh_token"

	// Password represents the OAuth 2.0 password grant
	Password Grant = "password"
)

// ListOAuthClients lists all of the users OAuth clients
func (a *Account) ListOAuthClients() (clients []*OAuthClient, err error) {
	err = a.get(a.auth.WithScope(scope.Clients), "/api/v2/clients", &clients)
	return clients, err
}

// FindOAuthClient gets the OAuth client with the name
func (a *Account) FindOAuthClient(name string) (client *OAuthClient, err error) {
	err = a.get(a.auth.WithScope(scope.Clients), fmt.Sprintf("/api/v2/clients/%s", name), &client)
	return client, err
}

// CreateOAuthClient creates a new OAuth client
func (a *Account) CreateOAuthClient(client *OAuthClient) (c *OAuthClient, err error) {
	err = a.post(a.auth.WithScope(scope.Clients), "/api/v2/clients", client, &c)
	return c, err
}

// EditOAuthClient updates the OAuth client
func (a *Account) EditOAuthClient(name string, edits *OAuthClient) (err error) {
	return a.patch(a.auth.WithScope(scope.Clients), fmt.Sprintf("/api/v2/clients/%s", name), edits, nil)
}

// RemoveOAuthClient removes the OAuth client
func (a *Account) RemoveOAuthClient(name string) (err error) {
	return a.del(a.auth.WithScope(scope.Clients), fmt.Sprintf("/api/v2/clients/%s", name))
}
