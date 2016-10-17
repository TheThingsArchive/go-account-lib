// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package test

import "github.com/TheThingsNetwork/go-account-lib/tokenkey"

// Issuer is an issuer id used to build the test tokens
var Issuer = "test-issuer"

// PrivateKey is a private key that is used to sign the tesk tokens
var PrivateKey = []byte(`
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
-----END RSA PRIVATE KEY-----`)

// PublicKey is the public key associated with PrivateKey
var PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzxNptNCOwLZM5SFcUI/l
pqoN2YofNAFN4jNyv1L1qHb+6V82+pRA5QnGJCAOo3bcx3Rra170hJfJUE4f3YxC
4OaCFflzIwZcG5gaSl25Bm4ISO/y8HbOUtDFl8djkvcQEDu1fndKy7Hfmvghn+Rn
V68YLfXocmjM1XFD6aiRQhWDU4Pkdhso+hMd56qPziLo0X+Ebt9GuDj3VsfxDQz/
li7o2RLj7IRfX+mMn1K5OFOsSm0TuMxhaIhpxizoh3XkW7oe6H66uvoEVkI/0rBe
KeWzmJIt+gGbRnZnqb3tevXnSYKh9aQb/o5VYfC6x6xBpXA2yIQQILheNHuwn7t/
aQIDAQAB
-----END PUBLIC KEY-----`)

// Provider is a tokenkey.Provider that always returns PublicKey
var Provider = tokenkey.ConstProvider(string(PublicKey), "RS256")
