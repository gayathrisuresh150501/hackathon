Atreugo
=======

[![Build Status](https://travis-ci.org/savsgio/atreugo.svg?branch=master)](https://travis-ci.org/savsgio/atreugo)
[![Coverage Status](https://coveralls.io/repos/github/savsgio/atreugo/badge.svg?branch=master)](https://coveralls.io/github/savsgio/atreugo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/savsgio/atreugo)](https://goreportcard.com/report/github.com/savsgio/atreugo)
[![GoDoc](https://godoc.org/github.com/savsgio/atreugo?status.svg)](https://godoc.org/github.com/savsgio/atreugo)
[![GitHub release](https://img.shields.io/github/release/savsgio/atreugo.svg)](https://github.com/savsgio/atreugo/releases)

Micro-framework to make simple the use of routing and middlewares in [fasthttp](https://github.com/valyala/fasthttp).

***The project use [dep](https://golang.github.io/dep/) manager dependencies.***


## Note:
`*atreugo.RequestCtx` is equal than `*fasthttp.RequestCtx`, but adding extra funtionality, so you can use
the same functions of `*fasthttp.RequestCtx`. Don't worry :smile:

## Example:

```go
package main

import (
	"errors"

	"github.com/valyala/fasthttp"
	"github.com/savsgio/atreugo"
)

func main() {
	config := &atreugo.Config{
		Host: "0.0.0.0",
		Port: 8000,
	}
	server := atreugo.New(config)

	fnMiddlewareOne := func(ctx *atreugo.RequestCtx) (int, error) {
		return fasthttp.StatusOK, nil
	}

	fnMiddlewareTwo := func(ctx *atreugo.RequestCtx) (int, error) {
		// Disable this middleware if you don't want to see this error
		return fasthttp.StatusBadRequest, errors.New("Error example")
	}

	server.UseMiddleware(fnMiddlewareOne, fnMiddlewareTwo)

	server.Path("GET", "/", func(ctx *atreugo.RequestCtx) error {
		return ctx.HTTPResponse("<h1>Atreugo Micro-Framework</h1>")
	})

	server.Path("GET", "/jsonPage", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Atreugo": true})
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

```

## Useful third-party libraries

- [session](https://github.com/fasthttp/session)
- [websocket](https://github.com/fasthttp/websocket)

Contributing
============

**Feel free to contribute it or fork me...** :wink:
