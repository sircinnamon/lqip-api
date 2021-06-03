package server

import (
	"github.com/valyala/fasthttp"
	"imagehandler"
	"log"
	"argstructs"
	"fmt"
)

func Hw() {
	log.Println("Hello from Server")
	imagehandler.Hw()
}

func testEndpointHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Content-Type", "image/svg+xml")
	log.Println("Starting image conversion...")
	ctx.SetBody([]byte(imagehandler.TestRun(imgArgs)))
	log.Println("Done!")
}

func syncPostHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	log.Println("Starting image conversion...")
	svg, err := imagehandler.SyncRun(imgArgs, ctx.PostBody())
	if err != nil {
		log.Fatal(err)
		ctx.Error("Conversion Failed", fasthttp.StatusInternalServerError)
		return
	}
	log.Println("Done!")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Content-Type", "image/svg+xml")
	ctx.SetBody([]byte(svg))
}

func ListenAndServe(args *argstructs.ServerArgs, imgArgs *argstructs.ImageHandlerArgs) {
	router := func(ctx *fasthttp.RequestCtx){
		log.Println(string(ctx.Path()))
		switch string(ctx.Path()) {
		case "/":
			switch string(ctx.Method()){
			case "POST":
				syncPostHandler(imgArgs, ctx)
			default:
				testEndpointHandler(imgArgs, ctx)
			}
		case "/test":
			testEndpointHandler(imgArgs, ctx)
		default:
			ctx.Error("Not Found", fasthttp.StatusNotFound)
		}
	}

	listenHost := fmt.Sprintf(":%d", args.Port)

	if err := fasthttp.ListenAndServe(listenHost, router); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}