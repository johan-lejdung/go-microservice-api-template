# go-microservice-api-template

[![CircleCI](https://circleci.com/gh/johan-lejdung/go-microservice-api-template.svg?style=svg)](https://circleci.com/gh/johan-lejdung/go-microservice-api-template)
[![codecov](https://codecov.io/gh/johan-lejdung/go-microservice-api-template/branch/master/graph/badge.svg)](https://codecov.io/gh/johan-lejdung/go-microservice-api-template)

A template for a go microservice with a REST-API.

For a service configured with PubSub look at https://github.com/johan-lejdung/go-microservice-pubsub-template

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

Open up `.env` and change the settings necessary.

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
go generate ./....
```

Run tests
```
go test ./...
```

# Code Guide

## Api
The api is found in the `goservice` folder and has the service it uses injected to it's structure.

To have multiple endpoints, simply add endpoints in `InitAPIRoute()`.

Middleware specific to endpoint can be added on line 25 in `api.go`:

```
Handler(negroni.New(
    potentialMiddleware,
    negroni.HandlerFunc(a.GoEndpoint()),
))
```

Currently there is one POST endpoint and one GET endpoint.

**POST endpoint**
```
POST localhost:8080/endpoint/
```
It expects a body on this form
```
{
    value: "someValue"
}
```

Try it out with
```
CURL -v -d '{"value":"this is a value"}' localhost:8080/endpoint/
```

**GET endpoint**
```
GET localhost:8080/endpoint/{id}
```

Try it out with
```
CURL -v localhost:8080/endpoint/1
```

## Database
The database will automatically apply migrations if the variable in `.env` called `ENV` is either `dev` or `test`.

I am using https://github.com/mattes/migrate for database migrations.

Run a local database in docker with `docker-compose up`.

## Bootstrap
I am using https://github.com/facebookgo/inject for dependency injection.

By injecting implementations of interfaces in the `bootstrapApp` you can easily inject them in structs such as:

```
type TestStruct struct {
    Variable pkg.InterfaceType `inject:""`
}
```

## Vendor folder
I have the vendor folder checked into the repo, for reproducibility.
