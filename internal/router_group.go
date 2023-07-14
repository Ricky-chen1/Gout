package gout

type RouterGroup struct {
	prefix     string
	middleware []HandlerFunc
	engine     *Engine
	parent     *RouterGroup //支持嵌套路由组
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	e := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: e,
		parent: group,
	}

	e.Groups = append(e.Groups, newGroup)
	return newGroup
}

//路由分组注册路由
func (group *RouterGroup) addRouter(method string, left string, handler HandlerFunc) {
	path := group.prefix + left

	group.engine.router.addRouter(method, path, handler)
}

func (group *RouterGroup) Use(mid ...HandlerFunc) {
	group.middleware = append(group.middleware, mid...)
}

func (group *RouterGroup) GET(left string, handler HandlerFunc) {
	group.addRouter("GET", left, handler)
}

func (group *RouterGroup) POST(left string, handler HandlerFunc) {
	group.addRouter("POST", left, handler)
}

func (group *RouterGroup) PUT(left string, handler HandlerFunc) {
	group.addRouter("PUT", left, handler)
}

func (group *RouterGroup) DELETE(left string, handler HandlerFunc) {
	group.addRouter("DELETE", left, handler)
}
