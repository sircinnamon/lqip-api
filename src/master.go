package main

import (
	"server"
	"log"
	"argstructs"
	"github.com/spf13/pflag"
)

var servArgs argstructs.ServerArgs
var imgArgs argstructs.ImageHandlerArgs

func init() {
	// Init flags
	pflag.IntVarP(&servArgs.Port, "port", "p", 80, "Port to listen on")

	pflag.BoolVarP(&servArgs.AllowAsync, "async", "a", false, "Allow drop-off/pickup async requests")
	pflag.IntVar(&servArgs.AsyncCacheExpiry, "cacheLife", 600, "Time (in seconds) an async svg should be stored")
	pflag.IntVar(&servArgs.AsyncCacheGC, "cacheGC", 900, "Cadence (in seconds) in between scraping cache for expired content")
	
	pflag.BoolVarP(&servArgs.AllowPostback, "postback", "r", false, "Allow drop-off/send-back async requests")

	pflag.IntVarP(&imgArgs.Shapes, "defaultShapeCount", "s", 16, "Default number of shapes in an LQIP")
	pflag.BoolVar(&imgArgs.AllowShapeCountQP, "allowShapeCountQP", true, "Allow user to specify non-default shape count")
	pflag.IntVar(&imgArgs.MaxShapeCountQP, "maxShapeCountQP", 32, "Maximum user shape count specifiable")

	pflag.IntVarP(&imgArgs.Mode, "defaultMode", "m", 1, "Default type of shape to generate for the LQIP [0-8]")
	pflag.StringVar(&imgArgs.AllowedModeQPs, "allowedModeQPs", "12345678", "Allowable modes specifiable by user")

	pflag.IntVarP(&imgArgs.Blur, "blur", "b", 12, "Default level of Gaussian blur filter")
	pflag.BoolVar(&imgArgs.AllowBlurQP, "allowBlurQP", true, "Allow user to specify different blur level")


	pflag.Parse()
}

func main() {
	log.Println("Running")
	// server.Hw()
	server.ListenAndServe(&servArgs, &imgArgs)
}