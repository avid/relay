package relay

import (
	"github.com/valyala/fasthttp"
)

func fork(ctx *fasthttp.RequestCtx, queue []Middleware) *runner {
	impl := new(runner)
	impl.ctx = ctx
	impl.queue = queue
	impl.step = 0
	impl.size = len(queue)
	return impl
}

type runner struct {
	ctx *fasthttp.RequestCtx
	step int
	size int
	queue []Middleware
}

func (r *runner) invoke() error {
	if r.step>=r.size {
		return nil
	}
	fn := r.queue[r.step]
	r.step++
	return fn(r.ctx, r.invoke)
}

