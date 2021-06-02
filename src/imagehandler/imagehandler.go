package imagehandler

import (
	"log"
	"fmt"
	"github.com/sircinnamon/sqip"
)

func Hw(){
	fmt.Println("Hello from Image Handler")
}
func Run(){
	in := "testimg.jpg"
	svg, w, h, err := sqip.Run(in, 256, 8, 1, 128, 0, 1, "")

	if err != nil {
		log.Fatal(err)
	}

	out := "tmp"
	if err := sqip.SaveFile(out, svg); err != nil {
		log.Fatal(err)
	}

	tag := sqip.ImageTag(out, sqip.Base64(svg), w, h)
	log.Print(tag)
}