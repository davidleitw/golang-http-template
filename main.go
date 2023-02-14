package main

import (
	"handler/function"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	r.POST("/", function.Handler)
	log.Fatalln(fasthttp.ListenAndServe(":8080", r.Handler))
}
