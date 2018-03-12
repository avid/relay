# relay
Simple golang middleware for https://github.com/valyala/fasthttp

### notes
Be careful, middlewares should be stateless.

### example

code:
```golang
func main() {
	handler := relay.New()

	handler.Use(doSomething1)
	handler.Use(doSomething2)
	handler.Use(doSomething3)

	fasthttp.ListenAndServe("127.0.0.1:80", handler.Handle)
	c := make(chan bool)
	<-c
}

func doSomething1 (ctx *fasthttp.RequestCtx, next relay.Next) error {
	fmt.Println("in 1")
	err := next()
	fmt.Println("out 1")
	return err
}

func doSomething2 (ctx *fasthttp.RequestCtx, next relay.Next) error {
	fmt.Println("in 2")
	err := next()
	fmt.Println("out 2")
	return err
}

func doSomething3 (ctx *fasthttp.RequestCtx, next relay.Next) error {
	fmt.Println("in 3")
	err := next()
	fmt.Println("out 3")
	return err
}
```

output:
```
in 1
in 2
in 3
out 3
out 2
out 1
```
