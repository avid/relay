package relay

import (
	"github.com/valyala/fasthttp"
)

type Relay interface {
	Push(mws ...Middleware)
	Handle(ctx *fasthttp.RequestCtx)
}

func New() Relay {
	impl := new(relayImpl)
	impl.queue = make([]Middleware,0)
	return impl
}

type relayImpl struct {
	queue []Middleware
}

func (r *relayImpl) Push(mws ...Middleware) {
	r.queue = append(r.queue, mws...)
}

func (r *relayImpl) Handle(ctx *fasthttp.RequestCtx) {
	runner := fork(ctx, r.queue)
	err := runner.invoke()
	if err != nil {
		ctx.SetContentType("text/plain")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
	}
}
