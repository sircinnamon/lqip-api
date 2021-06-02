package imagehandler

import (
	"log"
	"runtime"
	"github.com/sircinnamon/sqip"
	"argstructs"
	"bytes"
	"image"
	_ "image/gif"  // register gif
	_ "image/jpeg" // register jpeg
	_ "image/png"  // register png
)

func Hw(){
	log.Println("Hello from Image Handler")
}

func decodeBody(body []byte) (img image.Image, err error){
	// log.Println("About to decode")
	// log.Printf("Body len = %d\n", len(body))
	bodyReader := bytes.NewReader(body)
	img, _, err = image.Decode(bodyReader)
	// log.Println("decoded")
	return img, err
}

func TestRun(args *argstructs.ImageHandlerArgs) string{
	in := "testimg.jpg"
	workers := runtime.NumCPU()
	svg, _, _, err := sqip.Run(in, 256, args.Shapes, 1, 128, 0, workers, "")

	if err != nil {
		log.Fatal(err)
	}

	return svg
}

func SyncRun(args *argstructs.ImageHandlerArgs, body []byte) (svg string, err error){
	img, err := decodeBody(body)
	if err != nil {
		return "", err
	}
	workers := runtime.NumCPU()
	svg, _, _, err = sqip.RunLoaded(img, 256, args.Shapes, 1, 128, 0, workers, "")
	// log.Println(svg)
	// log.Println(err)

	return svg, err
}