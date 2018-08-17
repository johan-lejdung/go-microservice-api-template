# go-microservice-api-template

A template for a go microservice with a API.

Features:
- Dependency injection
- Database connection
- Database migrations
- Fluentd logging
- Api setup including tests
- Kubernetes compatible logging/liveness endpoints
- CricleCI + CodeCov configuration files

# Make it yours!

## Rename

Start by forking this repository or just download the code.

Replace all occourances of `johan-lejdung` with whatever your github username/organization is.

Replace all occourances of `go-microservice-api-template` with whatever you are calling your project.

## Env Config

Open up `.env` and change the settings nessicary.

## CircleCI and CodeCov

Simple register an account at CircleCI and CodeCov, copy the CodeCov token to the environment variables on CircleCI like so:

`CODECOV_TOKEN` = `<TOKEN>`

# Installation
Add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    export GOPATH=/Users/username/go

    export PATH=$GOPATH/bin:$PATH

Where username is the username of your profile.

Then install dep:

```
brew install dep
```

Run this command in the correct folder:

```
dep ensure
```

# Testing

Install mockery

```
go get github.com/vektra/mockery/.../
```

Generate files
```
go generate ./services/....
```

Run tests
```
go test ./services/...
```

# Code Guide

## Api
The api is found in the `api folder` and has the service it uses injected to it's structure.

To have multiple endpoints, simply add endpoints in `InitAPIRoute()`.

Middleware specific to endpoint can be added on line 25 in `api.go`:

```
Handler(negroni.New(
    potentialMiddleware,
    negroni.HandlerFunc(a.GoEndpoint()),
))
```

## Database
The database will automatically apply migrations if the variable in `.env` called `ENV` is either `dev` or `test`.

I am using https://github.com/mattes/migrate for database migrations.

## Bootstrap
I am using https://github.com/facebookgo/inject for dependency injection.

By injecting implementations of interfaces in the `bootstrapApp` you can easily inject them in structs such as:

```
type TestStruct struct {
    Variable pkg.InterfaceType `inject:""`
}
```
