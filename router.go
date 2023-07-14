package gout

import (
	"strings"
)

type Router struct {
	roots map[string]*node
	trees map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots: make(map[string]*node),
		trees: make(map[string]HandlerFunc),
	}
}

// 请求处理逻辑
func (r *Router) Handle(ctx *Context) {
	node, params := r.getRouter(ctx.Request.Method, ctx.Request.URL.Path)

	if node != nil {
		//找到后调用该路由处理函数
		ctx.Params = params //将解析出的路由参数放入上下文中

		key := ctx.Request.Method + "-" + node.path
		handler := r.trees[key]
		//用户自定义请求处理函数放入上下文中
		ctx.Handlers = append(ctx.Handlers, handler)
	} else {
		ctx.Handlers = append(ctx.Handlers, func(ctx *Context) {
			ctx.String(404, "404 NOT FOUND: %s\n", ctx.Request.URL.Path)
		})
	}

	ctx.Next()
}

func parsePath(path string) []string {
	items := strings.Split(path, "/")

	parts := make([]string, 0)
	for _, item := range items {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 注册路由
func (r *Router) addRouter(method string, path string, handler HandlerFunc) {
	parts := parsePath(path)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	key := method + "-" + path

	r.roots[method].insert(path, parts, 0)
	r.trees[key] = handler
}

// 获取匹配路由
func (r *Router) getRouter(method string, path string) (*node, map[string]string) {
	foundParts := parsePath(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	//该方法字典树未找到
	if !ok {
		return nil, nil
	}

	n := root.find(foundParts, 0)

	if n != nil {
		parts := parsePath(n.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = foundParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(foundParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
