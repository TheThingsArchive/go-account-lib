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
- `oauth`: A wrapper around `golang.org/x/oauth2` that provides extra
  functionality like caching, following redirects etc.

# Contribute

**Of course follow the coding golang guidelines**

## Be clear and consistent

1. Do not use shortened form anywhere.
2. Type-scoped functions. Functions do one thing on one type.
3. Type functions names and requirement
 - Types collection functions should have like: `ListTypes`, `StreamTypes`.
 - Type must have the these function: `RegisterType`, `EditType`, `DeleteType`, `GetType`.
 - Type attribute functions should be under this form: `AddTypeAttribute`, `RemoveTypeAttribute`, `ListTypeAttribute`, `GetTypeAttribute`.
4. Other functions form are allowed as long as they are justified and start with verb and refer the type, ex: `VerbTypeDoSmth`.
5. Documents your codes. This a librabrary that will be use in numerous package.
6. Write tests.

## Check

Run `make test vet lint` (or separately) before pushing.
