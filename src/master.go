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

	pflag.IntVarP(&imgArgs.Shapes, "defaultShapeCount", "s", 16, "Default number of shapes in an LQIP")
	pflag.BoolVar(&imgArgs.AllowShapeCountQP, "allowShapeCountQP", true, "Allow user to specify non-default shape count")
	pflag.IntVar(&imgArgs.MaxShapeCountQP, "maxShapeCountQP", 32, "Maximum user shape count specifiable")

	pflag.Parse()
}

func main() {
	log.Println("Running")
	server.Hw()
	server.ListenAndServe(&servArgs, &imgArgs)
}