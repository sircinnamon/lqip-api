package main

import (
	"server"
	"fmt"
	"argstructs"
)

func init() {
	// Init
}

func main() {
	fmt.Println("Running")
	server.Hw()
	server.ListenAndServe(argstructs.ServerArgs{":9980"}, argstructs.ImageHandlerArgs{})
}