package server

import (
	"github.com/valyala/fasthttp"
	"imagehandler"
	"log"
	"argstructs"
)

func Hw() {
	log.Println("Hello from Server")
	imagehandler.Hw()
}

func testEndpointHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Content-Type", "image/svg+xml")
	log.Println("Starting image conversion...")
	ctx.SetBody([]byte(imagehandler.Run(imgArgs)))
	log.Println("Done!")
}

func ListenAndServe(args *argstructs.ServerArgs, imgArgs *argstructs.ImageHandlerArgs) {
	router := func(ctx *fasthttp.RequestCtx){
		log.Println(string(ctx.Path()))
		switch string(ctx.Path()) {
		case "/":
			testEndpointHandler(imgArgs, ctx)
		case "/test":
			testEndpointHandler(imgArgs, ctx)
		default:
			ctx.Error("Not Found", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(args.Port, router)
}