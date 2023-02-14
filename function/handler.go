package function

import (
	"bytes"
	"net/http"

	"github.com/valyala/fasthttp"
)

func Handler(ctx *fasthttp.RequestCtx) {
	buf := bytes.NewBufferString("Hello world ")
	buf.Write(ctx.Request.Body())
	ctx.SetBody(buf.Bytes())
	ctx.SetStatusCode(http.StatusOK)
}
