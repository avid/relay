package relay

import "github.com/valyala/fasthttp"

type Next func() error

type Middleware func(ctx *fasthttp.RequestCtx, next Next) error

