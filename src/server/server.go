package server

import (
	"github.com/valyala/fasthttp"
	"imagehandler"
	"fmt"
	"argstructs"
)

func Hw() {
	fmt.Println("Hello from Server")
	imagehandler.Hw()
}

func ListenAndServe(args argstructs.ServerArgs, imgArgs argstructs.ImageHandlerArgs) {
	m := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.Header.Set("Content-Type", "image/svg+xml")
		ctx.SetBody([]byte(imagehandler.Run(imgArgs)))
	}

	fasthttp.ListenAndServe(args.Port, m)
}