package server

import (
	"github.com/valyala/fasthttp"
	"imagehandler"
	"log"
	"argstructs"
	"fmt"
	"time"
	"github.com/patrickmn/go-cache"
	"github.com/google/uuid"
	"github.com/fasthttp/router"
)

var asyncStoreCache *cache.Cache

func InitCache(args *argstructs.ServerArgs){
	expiry := time.Duration(args.AsyncCacheExpiry)*time.Second
	garbage_collect := time.Duration(args.AsyncCacheGC)*time.Second
	asyncStoreCache = cache.New(expiry, garbage_collect)
}

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

func logReq(ctx *fasthttp.RequestCtx){
	switch string(ctx.Method()) {
	case "POST":
		log.Println(fmt.Sprintf("REQ %s %s - %+v - BODYSIZE %s", ctx.Method(), ctx.Path(), parseQP(ctx), ctx.Request.Header.Peek("Content-Length")))
	default:
		log.Println(fmt.Sprintf("REQ %s %s - %+v", ctx.Method(), ctx.Path(), parseQP(ctx)))
	}
}

func syncPostHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	logReq(ctx)
	log.Println("Starting image conversion...")
	post_body := ctx.PostBody()
	svg, err := imagehandler.SyncRun(imgArgs, &post_body, parseQP(ctx))
	if err != nil {
		log.Fatal(err)
		ctx.Error("Conversion Failed", fasthttp.StatusInternalServerError)
		return
	}
	log.Println("Done!")
	ctx.SetStatusCode(fasthttp.StatusAccepted)
	ctx.Response.Header.Set("Content-Type", "image/svg+xml")
	ctx.SetBody([]byte(svg))
}

func asyncPostHandler(imgArgs *argstructs.ImageHandlerArgs, ctx *fasthttp.RequestCtx) {
	logReq(ctx)
	log.Println("Starting image conversion...")
	post_body := ctx.PostBody()
	svgCh, errCh := imagehandler.AsyncRun(imgArgs, &post_body, parseQP(ctx))
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Content-Type", "image/svg+xml")
	token := uuid.New().String()
	ctx.SetBody([]byte(token))
	go func() {
		select {
		case svg := <- svgCh:
			log.Println("Done!")
			asyncStoreCache.Set(token, svg, cache.DefaultExpiration)
			log.Println(fmt.Sprintf("Cached at %s", token))
		case err := <- errCh:
			log.Fatal(err)
		}
	}()
}

func asyncGetHandler(ctx *fasthttp.RequestCtx, id string) {
	logReq(ctx)
	svg, found := asyncStoreCache.Get(id)
	// Maybe handle a "wait longer" case
	if found {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.Header.Set("Content-Type", "image/svg+xml")
		ctx.SetBody([]byte(svg.(string)))
	} else {
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}
}

func ListenAndServe(args *argstructs.ServerArgs, imgArgs *argstructs.ImageHandlerArgs) {
	r := router.New()
	r.POST("/", func(ctx *fasthttp.RequestCtx){syncPostHandler(imgArgs, ctx)})
	if(args.AllowAsync){
		r.POST("/async", func(ctx *fasthttp.RequestCtx){asyncPostHandler(imgArgs, ctx)})
		r.GET("/async/{id}", func(ctx *fasthttp.RequestCtx){asyncGetHandler(ctx, ctx.UserValue("id").(string))})
	}
	listenHost := fmt.Sprintf(":%d", args.Port)

	if err := fasthttp.ListenAndServe(listenHost, r.Handler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}