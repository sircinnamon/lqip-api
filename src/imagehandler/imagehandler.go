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
	"regexp"
	"strconv"
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

func Reblur(args *argstructs.ImageHandlerArgs, qps *argstructs.QueryParameters, svg string) string {
	blur := getBlur(args, qps)
	return reblur(svg, blur)
}

func reblur(svg string, blur int) string {
	blur_re := regexp.MustCompile(`(<feGaussianBlur stdDeviation=")(\d+)(")`)
	svg = blur_re.ReplaceAllString(svg, `${1}`+strconv.Itoa(blur)+`${3}`)
	return svg
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

func getBlur(args *argstructs.ImageHandlerArgs, qps *argstructs.QueryParameters) int{
	blur := args.Blur
	if(args.AllowBlurQP){
		if(qps.Blur > -1){	
			blur = qps.Blur
		}	
	}
	return blur
}

func SyncRun(args *argstructs.ImageHandlerArgs, body *[]byte, qps *argstructs.QueryParameters) (svg string, err error){
	img, err := decodeBody(body)
	if err != nil {
		return "", err
	}
	workers := runtime.NumCPU()
	shapecount := getShapeCount(args, qps)
	mode := getMode(args, qps)
	blur := getBlur(args, qps)
	// log.Println(shapecount)
	svg, _, _, err = sqip.RunLoaded(img, 256, shapecount, mode, 128, 0, workers, "")
	if(blur != 12){ // 12 is default in the library
		svg = reblur(svg, blur)
	}
	// log.Println(svg)
	// log.Println(err)

	return svg, err
}

func AsyncRun(args *argstructs.ImageHandlerArgs, body *[]byte, qps *argstructs.QueryParameters) (chan string, chan error){
	r := make(chan string)
	error := make(chan error)
	go func() {
		svg, syncErr := SyncRun(args, body, qps)
		if syncErr != nil {
			error <- syncErr
			return
		}
		r <- svg
	}()
	return r, error
}