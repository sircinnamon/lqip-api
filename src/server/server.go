package server

import (
	"github.com/valyala/fasthttp"
	"imagehandler"
	"fmt"
)

func Hw() {
	fmt.Println("Hello from Server")
	imagehandler.Hw()
	m := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.Header.Set("Content-Type", "image/svg+xml")
		ctx.SetBody([]byte(imagehandler.Run()))
	}

	fasthttp.ListenAndServe(":9980", m)
}