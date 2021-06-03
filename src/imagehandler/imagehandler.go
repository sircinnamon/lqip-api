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
	"strings"
	"fmt"
)

func Hw(){
	log.Println("Hello from Image Handler")
}

func decodeBody(body *[]byte) (img image.Image, err error){
	// log.Println("About to decode")
	// log.Printf("Body len = %d\n", len(body))
	bodyReader := bytes.NewReader(*body)
	img, _, err = image.Decode(bodyReader)
	// log.Println("decoded")
	return img, err
}

func getShapeCount(args *argstructs.ImageHandlerArgs, qps *argstructs.QueryParameters) int{
	count := args.Shapes
	// log.Println(args)
	if(args.AllowShapeCountQP){
		if(qps.Shapes > 0){
			count = qps.Shapes
			if(count > args.MaxShapeCountQP){
				count = args.MaxShapeCountQP
			}
		}
	}
	return count
}

func getMode(args *argstructs.ImageHandlerArgs, qps *argstructs.QueryParameters) int{
	mode := args.Mode
	if(qps.Mode > -1){	
		if(strings.Contains(args.AllowedModeQPs, fmt.Sprintf("%d", qps.Mode))){
			mode = qps.Mode
		}
	}
	return mode
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

func SyncRun(args *argstructs.ImageHandlerArgs, body *[]byte, qps *argstructs.QueryParameters) (svg string, err error){
	img, err := decodeBody(body)
	if err != nil {
		return "", err
	}
	workers := runtime.NumCPU()
	shapecount := getShapeCount(args, qps)
	mode := getMode(args, qps)
	// log.Println(shapecount)
	svg, _, _, err = sqip.RunLoaded(img, 256, shapecount, mode, 128, 0, workers, "")
	// log.Println(svg)
	// log.Println(err)

	return svg, err
}