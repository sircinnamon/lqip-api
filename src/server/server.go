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


func parseQP(ctx *fasthttp.RequestCtx) *argstructs.QueryParameters{
	var qps argstructs.QueryParameters

	// qps.X will be set to -1 if missing
	qps.Shapes, _ = ctx.QueryArgs().GetUint("shapecount")
	qps.Mode, _ = ctx.QueryArgs().GetUint("mode")

	return &qps
}

func syncPostHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	log.Println("Starting image conversion...")
	post_body := ctx.PostBody()
	svg, err := imagehandler.SyncRun(imgArgs, &post_body, parseQP(ctx))
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
		log.Println(fmt.Sprintf("REQ %s - %+v - BODYSIZE %s", ctx.Path(), parseQP(ctx), ctx.Request.Header.Peek("Content-Length")))
		switch string(ctx.Path()) {
		case "/":
			switch string(ctx.Method()){
			case "POST":
				syncPostHandler(imgArgs, ctx)
			default:
				ctx.Error("Method Not Allowed", fasthttp.StatusMethodNotAllowed)
			}
		default:
			ctx.Error("Not Found", fasthttp.StatusNotFound)
		}
	}

	listenHost := fmt.Sprintf(":%d", args.Port)

	if err := fasthttp.ListenAndServe(listenHost, router); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}