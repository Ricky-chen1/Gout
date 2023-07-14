package main

import (
	gout "gin_demo"
	"net/http"
)

func main() {
	r := gout.Default()
	r.GET("/ping", func(c *gout.Context) {
		c.JSON(http.StatusOK, gout.H{
			"msg": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
