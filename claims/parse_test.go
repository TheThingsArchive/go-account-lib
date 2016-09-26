// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

import (
	"reflect"
	"testing"
	"time"

	"github.com/TheThingsNetwork/go-account-lib/tokenkey"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/assertions"
)

const (
	privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAzxNptNCOwLZM5SFcUI/lpqoN2YofNAFN4jNyv1L1qHb+6V82
+pRA5QnGJCAOo3bcx3Rra170hJfJUE4f3YxC4OaCFflzIwZcG5gaSl25Bm4ISO/y
8HbOUtDFl8djkvcQEDu1fndKy7Hfmvghn+RnV68YLfXocmjM1XFD6aiRQhWDU4Pk
dhso+hMd56qPziLo0X+Ebt9GuDj3VsfxDQz/li7o2RLj7IRfX+mMn1K5OFOsSm0T
uMxhaIhpxizoh3XkW7oe6H66uvoEVkI/0rBeKeWzmJIt+gGbRnZnqb3tevXnSYKh
9aQb/o5VYfC6x6xBpXA2yIQQILheNHuwn7t/aQIDAQABAoIBAQCrhRztlEJqBZYz
xCo+4LIMFpdaNTobTWlBj/Pf3ct1OvtyOlfDvsDx9eKVUahOZcoBu8CuMvy+RyuM
xOlIDUHoH4ZoxTJFNKNeh+Je7rqvRLzADWBhJUdI+Xxxd8pWlSZNC+gNVKozhqX8
KsNPOVUQIAwbJbDf80aXFTZ3eBS5crIJuawIvqPLZ8/4Gd0QcWPjHmCIvJQuDTe7
WIzkhxoTCuTMUTDW6kMtjlishx/e11kFBMznXYxxgvL3KaKVoUhsMxbbAxpP3Kbe
fcY4Sb0NuFp0sRY0I0kFOwEqQeZ5mLHTfjLFVYwW/qIwit+c7gqFXsuDv0LbAQ7A
8KS8DyIBAoGBAPYs3OnN/zX8lt74mMleqC6V8iv26OLvgOS5OGjFyryoMYgyMB92
wjHGo3HTYi9UB47+3GCuj/2vTw2iUdwnVI6LxKV/bjSbwpDJQNs33HifWu8J6fLL
Nogj7UpgnKMZpoti6NNjhJRCKGGUqBCi4qvuRLx5cTO6hM/sY4E/4V/ZAoGBANdX
E67Bo1nE5VuM803/lj+FgZhYb5XHivrdBFq8dZinwWQ9XxfMr9n+1EBcFBhNyYCh
RgcO8LWFbVPK1bOgAOYTSTIwmJnNPVfEvqslQAlFjPlaGtpBSqUU0lA2O6PHHyCJ
IXu0XAAtrl1eCg4aiWANCAwmxIHrAKwFNNmd0fIRAoGBAPQ252VOxarSDP3f0vqZ
2/BzMo7o4HoZLV46XTqbVZe4p4K8fz8HenkU3RpToKjhDKqQLSIAqrn5S0x0Rg9I
OTs8bvXbqAGqr+cgsCWJkj9bn0NaK2uAq3V9Zq8NjvbCwJSwp9bleCX4R8UeS2hN
nt7/fdMYCvRNSepXURNswvFpAoGBAMcuQQN1Eq43BFtBLc+oqIYK7EtJCbWGA9R0
2NFA3pkcGjKo3at65fGC1zrMsL2mPcsf4VEoDZgpWW2XAUILrqkhj6O/9XbVs3ba
ge52Hxw0W+hM4uecWvoFH1+YOmQMC4uhq/nrYum7Vzv/ftd6zjSs+ROcTElLYKy8
iBz98LKxAoGBANZsW5ENFKUjO+V55nMtHsVNtvSZqeXxADWFf4qI90BP/zJJ/n46
QSQP2Icfoc5JZ3W6Lo3Dg1vvj8FOt2bsCZiB/tB/JG+LGpXVdbLxtLCCQOkwDQlg
tlK5XdKTB0/+UoxBOxxO2g5Jb7CCus35MPI4h2N8sZ/M5XKksMRBUzHB
-----END RSA PRIVATE KEY-----`

	publicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzxNptNCOwLZM5SFcUI/l
pqoN2YofNAFN4jNyv1L1qHb+6V82+pRA5QnGJCAOo3bcx3Rra170hJfJUE4f3YxC
4OaCFflzIwZcG5gaSl25Bm4ISO/y8HbOUtDFl8djkvcQEDu1fndKy7Hfmvghn+Rn
V68YLfXocmjM1XFD6aiRQhWDU4Pkdhso+hMd56qPziLo0X+Ebt9GuDj3VsfxDQz/
li7o2RLj7IRfX+mMn1K5OFOsSm0TuMxhaIhpxizoh3XkW7oe6H66uvoEVkI/0rBe
KeWzmJIt+gGbRnZnqb3tevXnSYKh9aQb/o5VYfC6x6xBpXA2yIQQILheNHuwn7t/
aQIDAQAB
-----END PUBLIC KEY-----`
)

// constProvider is a tokenkey Provider that always resturns the same tokenkey
type constProvider struct {
	key *tokenkey.TokenKey
}

func (c *constProvider) Get(server string, renew bool) (*tokenkey.TokenKey, error) {
	return c.key, nil
}

func (c *constProvider) Update() error {
	return nil
}

func ConstProvider(publicKey string) tokenkey.Provider {
	return &constProvider{
		key: &tokenkey.TokenKey{
			Algorithm: "RS256",
			Key:       publicKey,
		},
	}
}

func buildJWT(claims *Claims) string {
	builder := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		panic(err)
	}

	token, err := builder.SignedString(key)
	if err != nil {
		panic(err)
	}
	return token
}

func buildClaims(TTL time.Duration) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "foo",
			ExpiresAt: time.Now().Add(TTL).Unix(),
		},
		Email:    "example@example.com",
		Client:   "foo",
		Scope:    []string{"apps"},
		Type:     "handler",
		Username: "example",
		Name: struct {
			First string `json:"first"`
			Last  string `json:"last"`
		}{
			First: "John",
			Last:  "Doe",
		},
	}
}

func TestParseValidToken(t *testing.T) {
	a := New(t)

	provider := ConstProvider(publicKey)

	claims := buildClaims(time.Hour)

	token := buildJWT(claims)

	parsed, err := FromToken(provider, token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}

func TestParseExpired(t *testing.T) {
	a := New(t)

	provider := ConstProvider(publicKey)

	claims := buildClaims(-1 * time.Hour)
	token := buildJWT(claims)

	_, err := FromToken(provider, token)
	a.So(err, ShouldNotBeNil)
}

func TestParseWithoutValidation(t *testing.T) {
	a := New(t)

	claims := buildClaims(time.Hour)
	token := buildJWT(claims)

	parsed, err := FromTokenWithoutValidation(token)
	a.So(err, ShouldBeNil)
	a.So(reflect.DeepEqual(parsed, claims), ShouldBeTrue)
}
