package imagehandler

import (
	"log"
	"runtime"
	"github.com/sircinnamon/sqip"
	"argstructs"
)

func Hw(){
	log.Println("Hello from Image Handler")
}

func Run(args *argstructs.ImageHandlerArgs) string{
	in := "testimg.jpg"
	workers := runtime.NumCPU()
	svg, _, _, err := sqip.Run(in, 256, args.Shapes, 1, 128, 0, workers, "")

	if err != nil {
		log.Fatal(err)
	}

	return svg
}