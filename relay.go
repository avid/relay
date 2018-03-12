package relay

import (
	"github.com/valyala/fasthttp"
)

// Relay represents common relay interface
type Relay interface {
	Use(mws ...Middleware)
	Handle(ctx *fasthttp.RequestCtx)
}

// New returns new relay instance
func New() Relay {
	impl := new(relayImpl)
	impl.queue = make([]Middleware,0)
	return impl
}

type relayImpl struct {
	queue []Middleware
}

func (r *relayImpl) Use(mws ...Middleware) {
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
