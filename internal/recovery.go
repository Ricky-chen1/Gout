package gout

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%s", err)
			log.Printf("%s\n", trace(msg))
			ctx.String(500, "Server Internal Error", nil)
		}
	}
}
