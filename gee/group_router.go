package gee

type RouterGroup struct {
	prefix string
	// middlewares []HandlerFunc
	// parent      *RouterGroup
	engine *Engine
}

func (group *RouterGroup) Group(prifix string) *RouterGroup {
	return &RouterGroup{
		prefix: group.prefix + prifix,
		engine: group.engine,
	}
}

func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
