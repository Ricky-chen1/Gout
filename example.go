package main

import (
	gout "gin_demo/internal"
	"net/http"
)

func main() {
	r := gout.Default()

	r.GET("/ping", func(ctx *gout.Context) {
		ctx.JSON(http.StatusOK, gout.H{
			"msg": "pong",
		})
	})
	r.Run(":8080")
}
