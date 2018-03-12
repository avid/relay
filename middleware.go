package relay

import "github.com/valyala/fasthttp"

// Next is a func to be called by Middleware
type Next func() error

// Middleware represents any func processing ctx
type Middleware func(ctx *fasthttp.RequestCtx, next Next) error

