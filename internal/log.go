package gout

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()

		log.Printf("[%d] %s in % v", ctx.StausCode, ctx.Request.URL.Path, time.Since(t))
	}
}
