package gout

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Request   *http.Request
	Writer    http.ResponseWriter
	Params    map[string]string
	Handlers  []HandlerFunc
	index     int
	StausCode int
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		Params:  make(map[string]string),
		index:   -1,
	}
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) JSON(code int, values interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.StausCode = code

	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(values); err != nil {
		http.Error(ctx.Writer, err.Error(), code)
	}
}

func (ctx *Context) String(code int, format string, values interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.StausCode = code

	ctx.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

func (ctx *Context) PostForm(key string) string {
	value := ctx.Request.FormValue(key)
	return value
}

func (ctx *Context) Query(key string) string {
	value := ctx.Request.URL.Query().Get(key)
	return value
}

func (ctx *Context) Param(key string) string {
	value := ctx.Params[key]
	return value
}

func (ctx *Context) Next() {
	ctx.index++
	l := len(ctx.Handlers)

	for ; ctx.index < l; ctx.index++ {
		ctx.Handlers[ctx.index](ctx)
	}
}
