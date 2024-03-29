module github.com/sircinnamon/lqip-api

require (
	argstructs v0.0.0
	github.com/fasthttp/router v1.3.14 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/fogleman/primitive v0.0.0-20200504002142-0373c216458b // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/sircinnamon/sqip v0.7.1-0.20190909152243-f9bf58f72d59 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/tdewolff/minify v2.3.6+incompatible // indirect
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
	golang.org/x/image v0.0.0-20210504121937-7319ad40d33e // indirect
	imagehandler v0.0.0
	server v0.0.0
)

replace imagehandler v0.0.0 => ./imagehandler

replace server v0.0.0 => ./server

replace argstructs v0.0.0 => ./argstructs
