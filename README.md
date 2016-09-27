[![GoDoc](https://godoc.org/github.com/TheThingsNetwork/go-account-lib?status.svg)](https://godoc.org/github.com/TheThingsNetwork/go-account-lib)
# go-account-lib

This is a client-library for the TTN account server.

It provides a suite of packages that assist in requesting,
using and validating access tokens.

It consists of these packages:

- `account`: A client package for actions on the account server
- `auth`: A library of strategies that decorate http request to authorize them
- `cache`: A library of caching strategies for storing and retrieving tokens
- `claims`: Helpers to parse and validate access tokens as well as check their
  scopes and access rights.
- `scope`: Contains all the scopes for an access token.
- `tokenkey`: A client library to fetch and cache token keys from the account
  server.
- `tokens`: A manager library that fetches and stores access tokens with
  different scopes based on an access token.
- `util`: An internal package that provides helpers for the other packages.

