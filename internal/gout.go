package gout

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	Groups []*RouterGroup
	router *Router
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.Groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

// 使用默认中间件
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.Groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}

	ctx := NewContext(r, w)
	ctx.Handlers = middlewares
	e.router.Handle(ctx)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.addRouter("POST", path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.router.addRouter("PUT", path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.router.addRouter("DELETE", path, handler)
}

// 启动Server监听该addr
func (e *Engine) Run(addr string) {
	//http路由转接入框架
	http.ListenAndServe(addr, e)
}
