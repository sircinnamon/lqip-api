package main

import (
	"server"
	"log"
	"argstructs"
)

func init() {
	// Init
}

func main() {
	log.Println("Running")
	server.Hw()
	servArgs := argstructs.ServerArgs{
		":9980",
	}
	imgArgs := argstructs.ImageHandlerArgs{
		16,
	}
	server.ListenAndServe(&servArgs, &imgArgs)
}